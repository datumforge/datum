# object

Config contains the configuration for the datum server


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**refreshInterval**|`integer`|RefreshInterval determines how often to reload the config<br/>||
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
|[**totp**](#totp)|`object`|||

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
|**shutdownGracePeriod**|`integer`|ShutdownGracePeriod sets the grace period for in flight requests before shutting down<br/>|no|
|**readTimeout**|`integer`|ReadTimeout sets the maximum duration for reading the entire request including the body<br/>|no|
|**writeTimeout**|`integer`|WriteTimeout sets the maximum duration before timing out writes of the response<br/>|no|
|**idleTimeout**|`integer`|IdleTimeout sets the maximum amount of time to wait for the next request when keep-alives are enabled<br/>|no|
|**readHeaderTimeout**|`integer`|ReadHeaderTimeout sets the amount of time allowed to read request headers<br/>|no|
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
|**certFile**|`string`|CertFile location for the TLS server<br/>||
|**certKey**|`string`|CertKey file location for the TLS server<br/>||
|**autoCert**|`boolean`|AutoCert generates the cert with letsencrypt, this does not work on localhost<br/>||

**Additional Properties:** not allowed  
<a name="servercors"></a>
### server\.cors: object

CORS settings for the server to allow cross origin requests


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|[**allowOrigins**](#servercorsalloworigins)|`string[]`|||
|**cookieInsecure**|`boolean`|CookieInsecure allows CSRF cookie to be sent to servers that the browser considers<br/>unsecured. Useful for cases where the connection is secured via VPN rather than<br/>HTTPS directly.<br/>||

**Additional Properties:** not allowed  
<a name="servercorsalloworigins"></a>
#### server\.cors\.allowOrigins: array

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
|[**supportedProviders**](#authsupportedproviders)|`string[]`||no|
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
|**refreshAudience**|`string`|RefreshAudience represents the audience for refreshing tokens.<br/>|no|
|**issuer**|`string`|Issuer represents the issuer of the tokens<br/>|yes|
|**accessDuration**|`integer`|AccessDuration represents the duration of the access token is valid for<br/>|no|
|**refreshDuration**|`integer`|RefreshDuration represents the duration of the refresh token is valid for<br/>|no|
|**refreshOverlap**|`integer`|RefreshOverlap represents the overlap time for a refresh and access token<br/>|no|
|**jwksEndpoint**|`string`|JWKSEndpoint represents the endpoint for the JSON Web Key Set<br/>|no|
|[**keys**](#authtokenkeys)|`object`||yes|

**Additional Properties:** not allowed  
<a name="authtokenkeys"></a>
#### auth\.token\.keys: object

**Additional Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|

<a name="authsupportedproviders"></a>
### auth\.supportedProviders: array

**Items**

**Item Type:** `string`  
<a name="authproviders"></a>
### auth\.providers: object

OauthProviderConfig represents the configuration for OAuth providers such as Github and Google


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**redirectUrl**|`string`|RedirectURL is the URL that the OAuth2 client will redirect to after authentication with datum<br/>||
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
|**clientId**|`string`|ClientID is the public identifier for the GitHub oauth2 client<br/>|yes|
|**clientSecret**|`string`|ClientSecret is the secret for the GitHub oauth2 client<br/>|yes|
|**clientEndpoint**|`string`|ClientEndpoint is the endpoint for the GitHub oauth2 client<br/>|no|
|[**scopes**](#authprovidersgithubscopes)|`string[]`||yes|
|**redirectUrl**|`string`|RedirectURL is the URL that the GitHub oauth2 client will redirect to after authentication with Github<br/>|yes|

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
|**clientId**|`string`|ClientID is the public identifier for the Google oauth2 client<br/>|yes|
|**clientSecret**|`string`|ClientSecret is the secret for the Google oauth2 client<br/>|yes|
|**clientEndpoint**|`string`|ClientEndpoint is the endpoint for the Google oauth2 client<br/>|no|
|[**scopes**](#authprovidersgooglescopes)|`string[]`||yes|
|**redirectUrl**|`string`|RedirectURL is the URL that the Google oauth2 client will redirect to after authentication with Google<br/>|yes|

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
|**displayName**|`string`|DisplayName is the site display name<br/>|yes|
|**relyingPartyId**|`string`|RelyingPartyID is the relying party identifier<br/>set to localhost for development, no port<br/>|yes|
|**requestOrigin**|`string`|RequestOrigin the origin domain for authentication requests<br/>include the scheme and port<br/>|yes|
|**maxDevices**|`integer`|MaxDevices is the maximum number of devices that can be associated with a user<br/>|no|
|**enforceTimeout**|`boolean`|EnforceTimeout at the Relying Party / Server. This means if enabled and the user takes too long that even if the browser does not<br/>enforce a timeout, the server will<br/>|no|
|**timeout**|`integer`|Timeout is the timeout in seconds<br/>|no|
|**debug**|`boolean`|Debug enables debug mode<br/>|no|

**Additional Properties:** not allowed  
<a name="authz"></a>
## authz: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|enables authorization checks with openFGA<br/>|no|
|**storeName**|`string`|name of openFGA store<br/>|no|
|**hostUrl**|`string`|host url with scheme of the openFGA API<br/>|yes|
|**storeId**|`string`|id of openFGA store<br/>|no|
|**modelId**|`string`|id of openFGA model<br/>|no|
|**createNewModel**|`boolean`|force create a new model<br/>|no|

**Additional Properties:** not allowed  
<a name="db"></a>
## db: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**debug**|`boolean`|debug enables printing the debug database logs<br/>|no|
|**databaseName**|`string`|the name of the database to use with otel tracing<br/>|no|
|**driverName**|`string`|sql driver name<br/>|no|
|**multiWrite**|`boolean`|enables writing to two databases simultaneously<br/>|no|
|**primaryDbSource**|`string`|dsn of the primary database<br/>|yes|
|**secondaryDbSource**|`string`|dsn of the secondary database if multi-write is enabled<br/>|no|
|**cacheTTL**|`integer`|cache results for subsequent requests<br/>|no|
|**runMigrations**|`boolean`|run migrations on startup<br/>|no|

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
|**dialTimeout**|`integer`|Dial timeout for establishing new connections, defaults to 5s<br/>||
|**readTimeout**|`integer`|Timeout for socket reads. If reached, commands will fail<br/>with a timeout instead of blocking. Supported values:<br/>  - `0` - default timeout (3 seconds).<br/>  - `-1` - no timeout (block indefinitely).<br/>  - `-2` - disables SetReadDeadline calls completely.<br/>||
|**writeTimeout**|`integer`|Timeout for socket writes. If reached, commands will fail<br/>with a timeout instead of blocking.  Supported values:<br/>  - `0` - default timeout (3 seconds).<br/>  - `-1` - no timeout (block indefinitely).<br/>  - `-2` - disables SetWriteDeadline calls completely.<br/>||
|**maxRetries**|`integer`|MaxRetries before giving up.<br/>Default is 3 retries; -1 (not 0) disables retries.<br/>||
|**minIdleConns**|`integer`|MinIdleConns is useful when establishing new connection is slow.<br/>Default is 0. the idle connections are not closed by default.<br/>||
|**maxIdleConns**|`integer`|Maximum number of idle connections.<br/>Default is 0. the idle connections are not closed by default.<br/>||
|**maxActiveConns**|`integer`|Maximum number of connections allocated by the pool at a given time.<br/>When zero, there is no limit on the number of connections in the pool.<br/>||

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
|**disableTimestamp**|`boolean`|DisableTimestamp disables the timestamp in the output<br/>||

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
|**sendGridApiKey**|`string`|SendGridAPIKey is the SendGrid API key to authenticate with the service<br/>||
|**fromEmail**|`string`|FromEmail is the default email to send from<br/>||
|**testing**|`boolean`|Testing is a bool flag to indicate we shouldn't be sending live emails, and instead should be writing out fixtures<br/>||
|**archive**|`string`|Archive is only supported in testing mode and is what is tied through the mock to write out fixtures<br/>||
|**datumListId**|`string`|DatumListID is the UUID SendGrid spits out when you create marketing lists<br/>||
|**adminEmail**|`string`|AdminEmail is an internal group email configured within datum for email testing and visibility<br/>||
|[**url**](#emailurl)|`object`|URLConfig for the datum registration<br/>||

**Additional Properties:** not allowed  
<a name="emailurl"></a>
### email\.url: object

URLConfig for the datum registration


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**base**|`string`|Base is the base URL used for URL links in emails<br/>||
|**verify**|`string`|Verify is the path to the verify endpoint used in verification emails<br/>||
|**invite**|`string`|Invite is the path to the invite endpoint used in invite emails<br/>||
|**reset**|`string`|Reset is the path to the reset endpoint used in password reset emails<br/>||

**Additional Properties:** not allowed  
<a name="sessions"></a>
## sessions: object

Config contains the configuration for the session store


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**signingKey**|`string`|SigningKey must be a 16, 32, or 64 character string used to encode the cookie<br/>||
|**encryptionKey**|`string`|EncryptionKey must be a 16, 32, or 64 character string used to encode the cookie<br/>||

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
|**enableTracing**|`boolean`|EnableTracing indicates whether tracing is enabled for the Sentry client<br/>||
|**traceSampler**|`number`|TracesSampler is the sampling rate for tracing in the Sentry client<br/>||
|**attachStacktrace**|`boolean`|AttachStacktrace indicates whether to attach stack traces to events in the Sentry client<br/>||
|**sampleRate**|`number`|SampleRate is the sampling rate for events in the Sentry client<br/>||
|**traceSampleRate**|`number`|TracesSampleRate is the sampling rate for tracing events in the Sentry client<br/>||
|**profileSampleRate**|`number`|ProfilesSampleRate is the sampling rate for profiling events in the Sentry client<br/>||
|**repanic**|`boolean`|Repanic indicates whether to repanic after capturing an event in the Sentry client<br/>||
|**debug**|`boolean`|Debug indicates whether debug mode is enabled for the Sentry client<br/>||
|**serverName**|`string`|ServerName is the name of the server running the Sentry client<br/>||

**Additional Properties:** not allowed  
<a name="posthog"></a>
## posthog: object

Config is the configuration for PostHog


**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled is a flag to enable or disable PostHog<br/>||
|**apiKey**|`string`|APIKey is the PostHog API Key<br/>||
|**host**|`string`|Host is the PostHog API Host<br/>||

**Additional Properties:** not allowed  
<a name="totp"></a>
## totp: object

**Properties**

|Name|Type|Description|Required|
|----|----|-----------|--------|
|**enabled**|`boolean`|Enabled is a flag to enable or disable the OTP service<br/>||
|**codeLength**|`integer`|CodeLength is the length of the OTP code<br/>||
|**issuer**|`string`|Issuer is the issuer for TOTP codes<br/>||
|**redis**|`boolean`|WithRedis configures the service with a redis client<br/>||
|**secret**|`string`|Secret stores a versioned secret key for cryptography functions<br/>||
|**recoveryCodeCount**|`integer`|RecoveryCodeCount is the number of recovery codes to generate<br/>||
|**recoveryCodeLength**|`integer`|RecoveryCodeLength is the length of a recovery code<br/>||

**Additional Properties:** not allowed  

