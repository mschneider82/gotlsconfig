

# gotlsconfig
`import "github.com/mschneider82/gotlsconfig"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package gotlsconfig makes it easy to get secure tlsconfig for testing and
development. It generates on the fly pub/private certificates.
This mitigates the usage of static private keys in go code.

Its better to use Self Singed Certificates instead of doing plain traffic!




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func New(subjectNames ...string) *tls.Config](#New)
* [func NewWithConfig(ssconfig SelfSignedConfig) (*tls.Config, error)](#NewWithConfig)
* [type KeyType](#KeyType)
* [type SelfSignedConfig](#SelfSignedConfig)


#### <a name="pkg-files">Package files</a>
[doc.go](https://github.com/mschneider82/gotlsconfig/doc.go) [gotlsconfig.go](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go) [selfsigned.go](https://github.com/mschneider82/gotlsconfig/selfsigned.go)


## <a name="pkg-constants">Constants</a>
``` go
const (
    EC256   = KeyType("P256")
    EC384   = KeyType("P384")
    RSA2048 = KeyType("2048")
    RSA4096 = KeyType("4096")
    RSA8192 = KeyType("8192")
)
```
Constants for all key types we support.




## <a name="New">func</a> [New](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go?s=147:191#L9)
``` go
func New(subjectNames ...string) *tls.Config
```
New generates a new selfsinged rsa4096 certificate for subjectNames with 10 years expiry



## <a name="NewWithConfig">func</a> [NewWithConfig](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go?s=567:633#L23)
``` go
func NewWithConfig(ssconfig SelfSignedConfig) (*tls.Config, error)
```
NewWithConfig gets a new tls.Config with custom settings




## <a name="KeyType">type</a> [KeyType](https://github.com/mschneider82/gotlsconfig/selfsigned.go?s=380:399#L27)
``` go
type KeyType string
```









## <a name="SelfSignedConfig">type</a> [SelfSignedConfig](https://github.com/mschneider82/gotlsconfig/selfsigned.go?s=459:610#L30)
``` go
type SelfSignedConfig struct {
    SAN          []string // Subject Alternative Names
    KeyType      KeyType
    Expire       time.Time
    Organization string
}

```
SelfSignedConfig configures a self-signed certificate.












