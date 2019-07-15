package gotlsconfig

import (
	"crypto/tls"
	"time"

	"github.com/xenolf/lego/certcrypto"
)

// New generates a new Selfsinged certificate with default settings
func New(names ...string) *tls.Config {
	if len(names) == 0 {
		names = []string{"localhost"}
	}
	ssconfig := SelfSignedConfig{
		SAN:     names,
		KeyType: certcrypto.RSA4096,
		Expire:  time.Now().Add(10 * time.Hour * 24 * 365),
	}
	cert, _ := newSelfSignedCertificate(ssconfig)
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

// NewWithConfig gets a new tls.Config with custom settings
func NewWithConfig(ssconfig SelfSignedConfig) (*tls.Config, error) {
	cert, err := newSelfSignedCertificate(ssconfig)
	return &tls.Config{Certificates: []tls.Certificate{cert}}, err
}
