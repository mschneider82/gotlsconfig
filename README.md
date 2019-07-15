

# gotlsconfig
`import "github.com/mschneider82/gotlsconfig"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func New(names ...string) *tls.Config](#New)
* [func NewWithConfig(ssconfig SelfSignedConfig) (*tls.Config, error)](#NewWithConfig)
* [type SelfSignedConfig](#SelfSignedConfig)


#### <a name="pkg-files">Package files</a>
[gotlsconfig.go](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go) [selfsigned.go](https://github.com/mschneider82/gotlsconfig/selfsigned.go)





## <a name="New">func</a> [New](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go?s=161:198#L11)
``` go
func New(names ...string) *tls.Config
```
New generates a new Selfsinged certificate with default settings



## <a name="NewWithConfig">func</a> [NewWithConfig](https://github.com/mschneider82/gotlsconfig/gotlsconfig.go?s=564:630#L25)
``` go
func NewWithConfig(ssconfig SelfSignedConfig) (*tls.Config, error)
```
NewWithConfig gets a new tls.Config with custom settings




## <a name="SelfSignedConfig">type</a> [SelfSignedConfig](https://github.com/mschneider82/gotlsconfig/selfsigned.go?s=287:384#L21)
``` go
type SelfSignedConfig struct {
    SAN     []string
    KeyType certcrypto.KeyType
    Expire  time.Time
}

```
SelfSignedConfig configures a self-signed certificate.













