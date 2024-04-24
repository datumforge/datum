# Kafka Broker

You can run the entire setup by running `./verify.sh`

## Individual components

This example demonstrates some basic operations with a Kafka broker proxied through Envoy.

Statistics collected by Envoy for the Kafka broker extension and related cluster metrics are also demonstrated.


### Start all containers

kafka/docker

```
docker compose pull
docker compose up --build -d
docker compose ps

           Name                      Command                State                            Ports
  -----------------------------------------------------------------------------------------------------------------------
  kafka_kafka-server_1   /etc/confluent/docker/run      Up             9092/tcp
  kafka_proxy_1          /docker-entrypoint.sh /usr ... Up             0.0.0.0:10000->10000/tcp, 0.0.0.0:8001->8001/tcp
  kafka_zookeeper_1      /etc/confluent/docker/run      Up (healthy)   2181/tcp, 2888/tcp, 3888/tcp
```

### Create a topic

Start by creating a Kafka topic with the name ``envoy-kafka-broker``:

```
export TOPIC="envoy-kafka-broker"
docker compose run --rm kafka-client kafka-topics --bootstrap-server proxy:10000 --create --topic $TOPIC
```

### Check topic

You can view the topics that Kafka is aware of with the ``kafka-topics --list`` argument.

Check that the topic you created exists:

```
docker compose run --rm kafka-client kafka-topics --bootstrap-server proxy:10000 --list | grep $TOPIC
```

### Send message using the producer

Next, send a message for the topic you have created using the ``kafka-console-producer``:

```
export MESSAGE="Welcome to Envoy and Kafka broker filter!"
docker compose run --rm kafka-client /bin/bash -c "echo $MESSAGE | kafka-console-producer --request-required-acks 1 --broker-list proxy:10000 --topic $TOPIC"
```

### Receive message using the consumer

Now you can receive the message using the `kafka-console-consumer` :

```
docker compose run --rm kafka-client kafka-console-consumer --bootstrap-server proxy:10000 --topic $TOPIC --from-beginning --max-messages 1 | grep "$MESSAGE"
```

### Check the broker stats

When you proxy to the Kafka broker, Envoy records various stats.

You can check the broker stats by querying the Envoy admin interface
(the numbers might differ a little as the kafka-client does not expose precise control over its network traffic):

```
curl -s "http://localhost:8001/stats?filter=kafka.kafka_broker" | grep -v ": 0" | grep "_request:"
  kafka.kafka_broker.request.api_versions_request: 9
  kafka.kafka_broker.request.create_topics_request: 1
  kafka.kafka_broker.request.fetch_request: 2
  kafka.kafka_broker.request.find_coordinator_request: 8
  kafka.kafka_broker.request.join_group_request: 2
  kafka.kafka_broker.request.leave_group_request: 1
  kafka.kafka_broker.request.list_offsets_request: 1
  kafka.kafka_broker.request.metadata_request: 12
  kafka.kafka_broker.request.offset_fetch_request: 1
  kafka.kafka_broker.request.produce_request: 1
  kafka.kafka_broker.request.sync_group_request: 1
```

### Check admin kafka cluster stats

Envoy also records cluster stats for the Kafka service:

```
curl -s "http://localhost:8001/stats?filter=cluster.kafka_service" | grep -v ": 0"
  cluster.kafka_service.max_host_weight: 1
  cluster.kafka_service.membership_healthy: 1
  cluster.kafka_service.membership_total: 1
```
