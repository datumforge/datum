# Datum on KraftCloud 

Spent some time testing out deploying the [datum](https://github.com/datumforge/datum) stack to KraftCloud, and for future Sarah, decided to write down some notes on the adventure

## Initial Learnings

Some learnings while getting started: 

1. You must use `-buildmode=pie` with `-extldflags -static-pie` when building the oci image otherwise you'll get this fun error:
    ```
    [    0.002614] ERR:  [appelfloader] datum: ELF executable is not position-independent!
    [    0.003306] ERR:  [appelfloader] datum: Parsing of ELF image failed: Exec format error (-8)
    ```
1. These files must be copied over from the `builder` in addition to the compiled server:
    ```
    COPY --from=builder /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/
    COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/
    ```
1. Exposed services must be running on `0.0.0.0` not `127.0.0.1` or `localhost`
1. All ports are exposed, so currently `EXPOSE` does not do anything
1. To expose a port to the internet, use `-p 443:17608` on the `deploy` command
1. Use the `scratch` image for the final base image

## Services 

The bare minimum services currently to run the [Datum](https://github.com/datumforge/datum) stack is with in memory-stores:
1. [Datum API](https://github.com/datumforge/datum)
1. [FGA](https://github.com/openfga/openfga)
1. Redis

To gain persistence, we also would need:

1. SQLite with persistent volume
1. Postgres for FGA persistence 

The following steps were timeboxed, so I went with the very minimum requirements of the stack to just get running (e.g. no persistence)

Besides updating configs, making sure the necessarily flags were exposed in `datum`, no code changes to the applications had to be made for either `datum` or `fga` to get it up and running. 

### Datum API

First up was setting up the datum unikernal oci image

#### Dockerfile

```Dockerfile
FROM --platform=linux/x86_64 golang:1.22.1-bookworm AS builder

WORKDIR /src

COPY . .

RUN set -xe; \
    CGO_ENABLED=1 GOOS=linux go build -a \
    -buildmode=pie \
    -ldflags "-linkmode external -extldflags -static-pie" \
    -o /datum \
    ;

FROM scratch

WORKDIR /home/nonroot

# Copy the binary from the builder
COPY --from=builder /datum /bin/datum
COPY --from=builder /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/

# Copy config 
COPY config/config-aio.example.yaml /config/.config.yaml

# Copy default model into image
COPY fga/model/datum.fga /fga/model/datum.fga

# Expose API port
EXPOSE 17608

ENTRYPOINT ["datum", "serve", "--debug", "--pretty"]
```

#### Kraftfile

This file is pretty basic, but it defines the `runtime` so you don't have you build your own base unikernal image, you can use `base:latest` with the choice of firecracker or qemu, this currently only has x86 support, so good luck running on arm if you don't want to build it yourself

```Kraftfile
spec: v0.6

runtime: base:latest

rootfs: ./Dockerfile.kc

cmd: ["datum", "serve", "--debug", "--pretty", "--config", "/config/.config.yaml"]
```

#### Deploy 

Export required environment variables

```bash
export KRAFTCLOUD_METRO=dal0
export KRAFTCLOUD_TOKEN=<REDACTED>
```

The following command will build and deploy the `datum` image and expose the API port to `443`

```bash
kraft cloud deploy -p 443:17608 . --memory 1000  -n datum
```

1. If you want a static name of the instance, use the `-n` flag to name it yourself. This is important for communicating on the private network 
1. Use a `service-group` if you want multiple services running within the same group, this wasn't necessary for this test deployment 
1. Expose the public port on `443` using `-p 443:<port>`
1. Make sure there is enough memory to run the service 


### FGA  

#### Dockerfile

Similar to the datum image, this is again a go-binary build so the `Dockerfile` will look very similar. 

1. Do not enable the `playground`, the services are unable to talk to each other on `0.0.0.0` internally so you will get timeouts and the instance will fail to start
1. Similarly, the `grpc` service needs to run on `127.0.0.1` instead of the default `0.0.0.0` in order to enable the internal communication 
1. The `http` service must run on `0.0.0.0` so the datum service can reach it at `fga.internal:8080`

```Dockerfile
FROM ghcr.io/grpc-ecosystem/grpc-health-probe:v0.4.25@sha256:6cc1dc0af87b35db2ca5fa9b1fbbc389e7570d8ad90ff84a54b6f7ac35cdb423 as grpc_health_probe
FROM --platform=linux/x86_64 golang:1.22.1-bookworm AS builder

WORKDIR /src

COPY . .

RUN set -xe; \
    CGO_ENABLED=0 GOOS=linux go build -a \
    -buildmode=pie \
    -ldflags "-extldflags -static-pie" \
    -o /bin/openfga ./cmd/openfga \
    ;

FROM scratch

WORKDIR /home/nonroot

# Copy FGA binary
COPY --from=builder /bin/openfga /openfga
COPY --from=grpc_health_probe /ko-app/grpc-health-probe /usr/local/bin/grpc_health_probe
COPY --from=builder /lib/x86_64-linux-gnu/libc.so.6 /lib/x86_64-linux-gnu/
COPY --from=builder /lib64/ld-linux-x86-64.so.2 /lib64/

# Expose FGA ports
EXPOSE 8081
EXPOSE 8080
EXPOSE 3000
EXPOSE 2112

ENTRYPOINT ["/openfga"]
```

#### Kraftfile

```Kraftfile
spec: v0.6

runtime: base:latest

rootfs: ./Dockerfile

cmd: ["/openfga", "run", "--experimentals", "check-query-cache", "--check-query-cache-enabled", "--log-format=json", "--grpc-addr=127.0.0.1:8081", "--playground-enabled=false"]
```

#### Deploy

```bash
kraft cloud deploy . --memory 512 -K Kraftfile -n fga
```

### Redis

This one was much easier because `Kraftcloud` provides an example repository that contains the required `Dockerfile` and `Kraftfile` so all you need follow their [docs](https://docs.kraft.cloud/guides/redis/)

#### Deploy 

Again, lets make sure we name it something consistent so we can configure `datum` with the right internal name to talk on the private network, `redis.internal` in this case:


```bash
 kraft cloud deploy -M 512 . -n redis
 ```

If you wanted `redis` exposed publicly, you could add the `tls` port to the command: 

```bash
 kraft cloud deploy -p 6379:6379/tls -M 512 . -n redis
 ```

## Instances

Now that we've deployed everything, we should be able to see our instances and we are off to the races:

```
cloud instance list                            
NAME   FQDN                                     STATE    CREATED AT      IMAGE                                                           MEMORY    ARGS                                                            BOOT TIME
datum  falling-meadow-cl81wzf3.dal0.kraft.host  running  10 seconds ago  manderson/datum@sha256:de9a76db6ffcd9e73d13510c1c5ff135ebce...  1000 MiB  datum serve --debug --pretty --config /config/.config.yaml      35.58 ms
redis                                           running  22 seconds ago  manderson/redis@sha256:9566df0f909b2f8640e2e692efabb469cf37...  512 MiB   /usr/bin/redis-server /etc/redis/redis.conf                     18.47 ms
fga                                             running  19 minutes ago  manderson/fga@sha256:de0fb2e2d96b61e4ffa5da47560117aa7ccf6a...  512 MiB   /openfga run --experimentals check-query-cache --check-quer...  25.18 ms
```

You'll notice there is no fqdn for `fga` or `redis`. This is because we do not need the authorization server or redis exposed to the public internet so we are using the internal fqdn for `datum --> fga` and `datum --> redis`. 

## Boot Times

Check out how fast these come up: 

```
kraft cloud  -p 443:17608 deploy manderson/datum:latest --memory 1000  -n datum
[+] deploying... done!                                                                                                                                                                                                                         [0.3s]

[●] Deployed successfully!
 │
 ├────────── name: datum                                                                                   
 ├────────── uuid: 786cf11d-713c-4450-9f46-609f47341747                                                    
 ├───────── state: running                                                                        
 ├─────────── url: https://billowing-brook-qiczhkki.dal0.kraft.host                                        
 ├───────── image: manderson/datum@sha256:de9a76db6ffcd9e73d13510c1c5ff135ebce356b773645ecbee44398b02b8d85 
 ├───── boot time: 34.93 ms        <----------- 30-40ms!!                                                                         
 ├──────── memory: 1000 MiB                                                                                
 ├─ service group: billowing-brook-qiczhkki                                                                
 ├── private fqdn: datum.internal                                                                          
 ├──── private ip: 172.16.110.5                                                                            
 └────────── args: datum serve --debug --pretty --config /config/.config.yaml  
 ``` 

 ```
 kraft cloud deploy manderson/fga:latest --memory 512  -n fga     
[+] deploying... done!                                                                                                                                                                                                                         [0.2s]

[●] Deployed successfully!
 │
 ├───────── name: fga                                                                                          
 ├───────── uuid: 7983ac13-5c94-4d8a-b061-beb4fa808694                                                         
 ├──────── state: running                                                                             
 ├──────── image: manderson/fga@sha256:de0fb2e2d96b61e4ffa5da47560117aa7ccf6aece32918485ca44e51ce3d7d88        
 ├──── boot time: 25.18 ms        <----- 20-30ms!!!                                                                              
 ├─────── memory: 512 MiB                                                                                      
 ├─ private fqdn: fga.internal                                                                                 
 ├─── private ip: 172.16.110.4                                                                                 
 └───────── args: /openfga run --experimentals check-query-cache --check-query-cache-enabled --log-format json 
```

## Follow-ups

1. The config is currently baked into the image, need to do some work to see if we can instead mount this as a volume, make it more configurable without rebuilding the image each time
1. Investigate `secrets` to use in the configs and the best way to provide these to the deployment. In the effort above, no secrets were added to the config, which means things that have externally dependencies like social logins, email sends, error reporting, etc were disabled. 
1. Test out the performance with a local sqlite deployed to Kraftcloud vs. using Turso in the same metro
1. Add `postgres` to the setup 
1. Test out `compose` to bring up all required services at once 
1. Publishing the images and portability of the images outside of Kraftcloud 