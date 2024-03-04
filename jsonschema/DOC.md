# object

Config contains the configuration for the datum server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**refresh\_interval**|`integer`|RefreshInterval determines how often to reload the config<br/>||
|[**server**](#server)|`object`|Server settings for the echo server<br/>|yes|
|[**auth**](#auth)|`object`|Auth settings including oauth2 providers and datum token configuration<br/>|yes|
|[**authz**](#authz)|`object`||yes|
|[**db**](#db)|`object`||yes|
|[**redis**](#redis)|`object`|Config for the redis client used to store key-value pairs<br/>||
|[**tracer**](#tracer)|`object`|Config defines the configuration settings for opentelemetry tracing<br/>||
|[**email**](#email)|`object`|Config for sending emails via SendGrid and managing marketing contacts<br/>||
|[**sessions**](#sessions)|`object`|Config contains the configuration for the session store<br/>||
|[**sentry**](#sentry)|`object`|Config settings for the Sentry client<br/>||
|[**posthog**](#posthog)|`object`|Config is the configuration for PostHog<br/>||

**Additional Properties:** not allowed  
<a name="server"></a>
## server: object

Server settings for the echo server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**debug**|`boolean`|Debug enables debug mode for the server<br/>|no|
|**dev**|`boolean`|Dev enables echo's dev mode options<br/>|no|
|**listen**|`string`|Listen sets the listen address to serve the echo server on<br/>|yes|
|**shutdown\_grace\_period**|`integer`|ShutdownGracePeriod sets the grace period for in flight requests before shutting down<br/>|no|
|**read\_timeout**|`integer`|ReadTimeout sets the maximum duration for reading the entire request including the body<br/>|no|
|**write\_timeout**|`integer`|WriteTimeout sets the maximum duration before timing out writes of the response<br/>|no|
|**idle\_timeout**|`integer`|IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled<br/>|no|
|**read\_header\_timeout**|`integer`|ReadHeaderTimeout sets the amount of time allowed to read request headers<br/>|no|
|[**tls**](#servertls)|`object`|TLS settings for the server for secure connections<br/>|no|
|[**cors**](#servercors)|`object`|CORS settings for the server to allow cross origin requests<br/>|no|

**Additional Properties:** not allowed  
<a name="servertls"></a>
### server\.tls: object

TLS settings for the server for secure connections


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled turns on TLS settings for the server<br/>||
|**cert\_file**|`string`|CertFile location for the TLS server<br/>||
|**cert\_key**|`string`|CertKey file location for the TLS server<br/>||
|**auto\_cert**|`boolean`|AutoCert generates the cert with letsencrypt, this does not work on localhost<br/>||

**Additional Properties:** not allowed  
<a name="servercors"></a>
### server\.cors: object

CORS settings for the server to allow cross origin requests


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|[**allow\_origins**](#servercorsallow_origins)|`string[]`|||
|**cookie\_insecure**|`boolean`|CookieInsecure allows CSRF cookie to be sent to servers that the browser considers<br/>unsecured. Useful for cases where the connection is secured via VPN rather than<br/>HTTPS directly.<br/>||

**Additional Properties:** not allowed  
<a name="servercorsallow_origins"></a>
#### server\.cors\.allow\_origins: array

**Items**

**Item Type:** `string`  
<a name="auth"></a>
## auth: object

Auth settings including oauth2 providers and datum token configuration


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled authentication on the server, not recommended to disable<br/>|no|
|[**token**](#authtoken)|`object`|Config defines the configuration settings for authentication tokens used in the server<br/>|yes|
|[**supported\_providers**](#authsupported_providers)|`string[]`||no|
|[**providers**](#authproviders)|`object`|OauthProviderConfig represents the configuration for OAuth providers such as Github and Google<br/>|no|

**Additional Properties:** not allowed  
<a name="authtoken"></a>
### auth\.token: object

Config defines the configuration settings for authentication tokens used in the server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**kid**|`string`|KID represents the Key ID used in the configuration.<br/>|yes|
|**audience**|`string`|Audience represents the target audience for the tokens.<br/>|yes|
|**refresh\_audience**|`string`|RefreshAudience represents the audience for refreshing tokens.<br/>|no|
|**issuer**|`string`|Issuer represents the issuer of the tokens<br/>|yes|
|**access\_duration**|`integer`|AccessDuration represents the duration of the access token is valid for<br/>|no|
|**refresh\_duration**|`integer`|RefreshDuration represents the duration of the refresh token is valid for<br/>|no|
|**refresh\_overlap**|`integer`|RefreshOverlap represents the overlap time for a refresh and access token<br/>|no|
|**jwks\_endpoint**|`string`|JWKSEndpoint represents the endpoint for the JSON Web Key Set<br/>|no|
|[**keys**](#authtokenkeys)|`object`||yes|

**Additional Properties:** not allowed  
<a name="authtokenkeys"></a>
#### auth\.token\.keys: object

**Additional Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|

<a name="authsupported_providers"></a>
### auth\.supported\_providers: array

**Items**

**Item Type:** `string`  
<a name="authproviders"></a>
### auth\.providers: object

OauthProviderConfig represents the configuration for OAuth providers such as Github and Google


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**redirect\_url**|`string`|RedirectURL is the URL that the OAuth2 client will redirect to after authentication with datum<br/>||
|[**github**](#authprovidersgithub)|`object`|ProviderConfig represents the configuration settings for a Github Oauth Provider<br/>|yes|
|[**google**](#authprovidersgoogle)|`object`|ProviderConfig represents the configuration settings for a Google Oauth Provider<br/>|yes|
|[**webauthn**](#authproviderswebauthn)|`object`|ProviderConfig represents the configuration settings for a Webauthn Provider<br/>|yes|

**Additional Properties:** not allowed  
<a name="authprovidersgithub"></a>
#### auth\.providers\.github: object

ProviderConfig represents the configuration settings for a Github Oauth Provider


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**client\_id**|`string`|ClientID is the public identifier for the GitHub oauth2 client<br/>|yes|
|**client\_secret**|`string`|ClientSecret is the secret for the GitHub oauth2 client<br/>|yes|
|**client\_endpoint**|`string`|ClientEndpoint is the endpoint for the GitHub oauth2 client<br/>|no|
|[**scopes**](#authprovidersgithubscopes)|`string[]`||yes|
|**redirect\_url**|`string`|RedirectURL is the URL that the GitHub oauth2 client will redirect to after authentication with Github<br/>|yes|

**Additional Properties:** not allowed  
<a name="authprovidersgithubscopes"></a>
##### auth\.providers\.github\.scopes: array

**Items**

**Item Type:** `string`  
<a name="authprovidersgoogle"></a>
#### auth\.providers\.google: object

ProviderConfig represents the configuration settings for a Google Oauth Provider


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**client\_id**|`string`|ClientID is the public identifier for the Google oauth2 client<br/>|yes|
|**client\_secret**|`string`|ClientSecret is the secret for the Google oauth2 client<br/>|yes|
|**client\_endpoint**|`string`|ClientEndpoint is the endpoint for the Google oauth2 client<br/>|no|
|[**scopes**](#authprovidersgooglescopes)|`string[]`||yes|
|**redirect\_url**|`string`|RedirectURL is the URL that the Google oauth2 client will redirect to after authentication with Google<br/>|yes|

**Additional Properties:** not allowed  
<a name="authprovidersgooglescopes"></a>
##### auth\.providers\.google\.scopes: array

**Items**

**Item Type:** `string`  
<a name="authproviderswebauthn"></a>
#### auth\.providers\.webauthn: object

ProviderConfig represents the configuration settings for a Webauthn Provider


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled is the provider enabled<br/>|no|
|**display\_name**|`string`|DisplayName is the site display name<br/>|yes|
|**relying\_party\_id**|`string`|RelyingPartyID is the relying party identifier<br/>|yes|
|**domain**|`string`|Domain is the domain of the site<br/>|yes|
|**request\_origin**|`string`|RequestOrigin the origin domain for authentication requests<br/>|yes|
|**max\_devices**|`integer`|MaxDevices is the maximum number of devices that can be associated with a user<br/>|yes|
|**enforce**|`boolean`|EnforceTimeout at the Relying Party / Server. This means if enabled and the user takes too long that even if the browser does not<br/>enforce a timeout, the server will<br/>|no|
|**timeout**|`integer`|Timeout is the timeout in seconds<br/>|no|
|**debug**|`boolean`|Debug enables debug mode<br/>|no|

**Additional Properties:** not allowed  
<a name="authz"></a>
## authz: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|enables authorization checks with openFGA<br/>|no|
|**store\_name**|`string`|name of openFGA store<br/>|no|
|**host\_url**|`string`|host url with scheme of the openFGA API<br/>|yes|
|**store\_id**|`string`|id of openFGA store<br/>|no|
|**model\_id**|`string`|id of openFGA model<br/>|no|
|**create\_new\_model**|`boolean`|force create a new model<br/>|no|

**Additional Properties:** not allowed  
<a name="db"></a>
## db: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**debug**|`boolean`|debug enables printing the debug database logs<br/>|no|
|**database\_name**|`string`|the name of the database to use with otel tracing<br/>|no|
|**driver\_name**|`string`|sql driver name<br/>|no|
|**multi\_write**|`boolean`|enables writing to two databases simultaneously<br/>|no|
|**primary\_db\_source**|`string`|dsn of the primary database<br/>|yes|
|**secondary\_db\_source**|`string`|dsn of the secondary database if multi-write is enabled<br/>|no|
|**cache\_ttl**|`integer`|cache results for subsequent requests<br/>|no|
|**run\_migrations**|`boolean`|run migrations on startup<br/>|no|

**Additional Properties:** not allowed  
<a name="redis"></a>
## redis: object

Config for the redis client used to store key-value pairs


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled to enable redis client in the server<br/>||
|**address**|`string`|Address is the host:port to connect to redis<br/>||
|**name**|`string`|Name of the connecting client<br/>||
|**username**|`string`|Username to connect to redis<br/>||
|**password**|`string`|Password, must match the password specified in the server configuration<br/>||
|**db**|`integer`|DB to be selected after connecting to the server, 0 uses the default<br/>||
|**dial\_timeout**|`integer`|Dial timeout for establishing new connections, defaults to 5s<br/>||
|**read\_timeout**|`integer`|Timeout for socket reads. If reached, commands will fail<br/>with a timeout instead of blocking. Supported values:<br/>  - `0` - default timeout (3 seconds).<br/>  - `-1` - no timeout (block indefinitely).<br/>  - `-2` - disables SetReadDeadline calls completely.<br/>||
|**write\_timeout**|`integer`|Timeout for socket writes. If reached, commands will fail<br/>with a timeout instead of blocking.  Supported values:<br/>  - `0` - default timeout (3 seconds).<br/>  - `-1` - no timeout (block indefinitely).<br/>  - `-2` - disables SetWriteDeadline calls completely.<br/>||
|**max\_retries**|`integer`|MaxRetries before giving up.<br/>Default is 3 retries; -1 (not 0) disables retries.<br/>||
|**min\_idle\_conns**|`integer`|MinIdleConns is useful when establishing new connection is slow.<br/>Default is 0. the idle connections are not closed by default.<br/>||
|**max\_idle\_conns**|`integer`|Maximum number of idle connections.<br/>Default is 0. the idle connections are not closed by default.<br/>||
|**max\_active\_conns**|`integer`|Maximum number of connections allocated by the pool at a given time.<br/>When zero, there is no limit on the number of connections in the pool.<br/>||

**Additional Properties:** not allowed  
<a name="tracer"></a>
## tracer: object

Config defines the configuration settings for opentelemetry tracing


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled to enable tracing<br/>||
|**provider**|`string`|Provider to use for tracing<br/>||
|**environment**|`string`|Environment to set for the service<br/>||
|[**stdout**](#tracerstdout)|`object`|StdOut settings for the stdout provider<br/>||
|[**otlp**](#tracerotlp)|`object`|OTLP settings for the otlp provider<br/>||

**Additional Properties:** not allowed  
<a name="tracerstdout"></a>
### tracer\.stdout: object

StdOut settings for the stdout provider


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**pretty**|`boolean`|Pretty enables pretty printing of the output<br/>||
|**disable\_timestamp**|`boolean`|DisableTimestamp disables the timestamp in the output<br/>||

**Additional Properties:** not allowed  
<a name="tracerotlp"></a>
### tracer\.otlp: object

OTLP settings for the otlp provider


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**endpoint**|`string`|Endpoint to send the traces to<br/>||
|**insecure**|`boolean`|Insecure to disable TLS<br/>||
|**certificate**|`string`|Certificate to use for TLS<br/>||
|[**headers**](#tracerotlpheaders)|`string[]`|||
|**compression**|`string`|Compression to use for the request<br/>||
|**timeout**|`integer`|Timeout for the request<br/>||

**Additional Properties:** not allowed  
<a name="tracerotlpheaders"></a>
#### tracer\.otlp\.headers: array

**Items**

**Item Type:** `string`  
<a name="email"></a>
## email: object

Config for sending emails via SendGrid and managing marketing contacts


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**send\_grid\_api\_key**|`string`|SendGridAPIKey is the SendGrid API key to authenticate with the service<br/>||
|**from\_email**|`string`|FromEmail is the default email to send from<br/>||
|**testing**|`boolean`|Testing is a bool flag to indicate we shouldn't be sending live emails, and instead should be writing out fixtures<br/>||
|**archive**|`string`|Archive is only supported in testing mode and is what is tied through the mock to write out fixtures<br/>||
|**datum\_list\_id**|`string`|DatumListID is the UUID SendGrid spits out when you create marketing lists<br/>||
|**admin\_email**|`string`|AdminEmail is an internal group email configured within datum for email testing and visibility<br/>||

**Additional Properties:** not allowed  
<a name="sessions"></a>
## sessions: object

Config contains the configuration for the session store


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**signing\_key**|`string`|SigningKey must be a 16, 32, or 64 character string used to encode the cookie<br/>||
|**encryption\_key**|`string`|EncryptionKey must be a 16, 32, or 64 character string used to encode the cookie<br/>||

**Additional Properties:** not allowed  
<a name="sentry"></a>
## sentry: object

Config settings for the Sentry client


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled indicates whether the Sentry client is enabled<br/>||
|**dsn**|`string`|DSN is the Data Source Name for the Sentry client<br/>||
|**environment**|`string`|Environment is the environment in which the Sentry client is running<br/>||
|**enable\_tracing**|`boolean`|EnableTracing indicates whether tracing is enabled for the Sentry client<br/>||
|**trace\_sampler**|`number`|TracesSampler is the sampling rate for tracing in the Sentry client<br/>||
|**attach\_stacktrace**|`boolean`|AttachStacktrace indicates whether to attach stack traces to events in the Sentry client<br/>||
|**sample\_rate**|`number`|SampleRate is the sampling rate for events in the Sentry client<br/>||
|**trace\_sample\_rate**|`number`|TracesSampleRate is the sampling rate for tracing events in the Sentry client<br/>||
|**profile\_sample\_rate**|`number`|ProfilesSampleRate is the sampling rate for profiling events in the Sentry client<br/>||
|**repanic**|`boolean`|Repanic indicates whether to repanic after capturing an event in the Sentry client<br/>||
|**debug**|`boolean`|Debug indicates whether debug mode is enabled for the Sentry client<br/>||
|**server\_name**|`string`|ServerName is the name of the server running the Sentry client<br/>||

**Additional Properties:** not allowed  
<a name="posthog"></a>
## posthog: object

Config is the configuration for PostHog


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled is a flag to enable or disable PostHog<br/>||
|**api\_key**|`string`|APIKey is the PostHog API Key<br/>||
|**host**|`string`|Host is the PostHog API Host<br/>||

**Additional Properties:** not allowed  

