# Gangly Configuration

Gangly reads a configuration file at startup. The path to the configuration file should be defined using the `--config` flag.

The configuration file should be in YAML format and contain a dictionary (alias hash or map) of key/value pairs. The available options are described below.

## Configuration Options

The following options can be set via the YAML configuration file.

### General Configuration

| Key                      | Description                                                                                                    |
|--------------------------|----------------------------------------------------------------------------------------------------------------|
| `host`                   | The address to listen on. Defaults to `0.0.0.0` (all interfaces).                                              |
| `port`                   | The port to listen on. Defaults to `8080`.                                                                     |
| `serveTLS`               | Should Gangly use TLS instead of plain HTTP? Defaults to `false`.                                              |
| `certFile`               | The public certificate file to use when using TLS. Defaults to `/etc/gangly/tls/tls.crt`.                      |
| `keyFile`                | The private key file when using TLS. Defaults to `/etc/gangly/tls/tls.key`.                                    |
| `trustedCAPath`          | Path to a root CA to trust for self-signed certificates at Oauth2 URLs.                                        |
| `httpPath`               | The path used by gangly to create URLs. Defaults to `""`.                                                      |
| `sessionSecurityKey`     | The session security key.                                                                                      |
| `sessionSalt`            | The session salt. Hardcoded default value `MkmfuPNHnZBBivy0L0aW`.                                              |
| `customHTMLTemplatesDir` | Path to a directory containing custom HTML templates.                                                          |
| `customAssetsDir`        | Path to a directory containing assets.                                                                         |

### Multi-Cluster Configuration

Multi-cluster configuration allows for specific configurations for each cluster within a single file.

| Key                   | Description                                                                                                    |
|-----------------------|----------------------------------------------------------------------------------------------------------------|
| `clusters`            | A dictionary of configurations by cluster name, with each cluster having its own set of options.              |

### Cluster-Specific Configuration

Each cluster can have the following configurations:

| Key                          | Description                                                                                                    |
|------------------------------|----------------------------------------------------------------------------------------------------------------|
| `clusterName`                | The name of the cluster. Used in the UI and the kubectl config instructions.                                   |
| `providerURL`                | OAuth2 provider URL. Must offer an endpoint `$providerURL/.well-known/openid-configuration` for discovery.     |
| `clientID`                   | API client ID as provided by the identity provider.                                                            |
| `clientSecret`               | API client secret as provided by the identity provider.                                                        |
| `allowEmptyClientSecret`     | Some identity providers accept an empty client secret, which is usually not a good idea. If you need to use an empty secret and accept the associated risks, then you can set this to `true`. Defaults to `false`.|
| `audience`                   | The endpoint that provides user profile information [optional]. Not required by all providers.                  |
| `scopes`                     | Used to specify the scope of the OAuth authorization request. Defaults to `["openid", "profile", "email", "offline_access"]`.|
| `redirectURL`                | Where to redirect after authentication. This should be a URL where Gangly is reachable. Typically, this must also be registered in the OAuth application with the OAuth provider.|
| `usernameClaim`              | The JWT claim to use as the username. This is used in the UI. Combined with the clusterName for the "user" part of kubeconfig. Defaults to `nickname`.|
| `emailClaim`                 | Deprecated. Defaults to `email`.                                                                               |
| `apiServerURL`               | The API server endpoint used for configuring kubectl.                                                          |
| `clusterCAPath`              | Path to find the CA bundle for the API server. Used for configuring kubectl. This path is typically mounted in the default location for workloads running on a Kubernetes cluster and usually doesn't need to be defined. Defaults to `/var/run/secrets/kubernetes.io/serviceaccount/ca.crt`.|
| `showClaims`                 | Display received claims. Defaults to `true`.                                                                   |

### Validation and Prefixes

Cluster-specific configurations can be validated and environment prefixes can be applied to allow configuration from environment variables.

## Additional Functions

- `NewMultiClusterConfig`: Creates a new multi-cluster configuration instance from a serialized configuration file.
- `Validate`: Verifies all properties of the configuration structure to ensure they are initialized.
- `GetRootPathPrefix`: Returns '/' if no prefix is specified, otherwise returns the configured path.
- `loadCerts`: Loads certificates for cluster configurations from specified paths.

## Use of Environment Variables

Environment variables can be used to override configurations specified in the YAML file by using the prefix `GANGLY` followed by the corresponding key name in uppercase and underscores for spaces.

Example: To override `clientSecret`, use the environment variable `GANGLY_CLIENT_SECRET`.
