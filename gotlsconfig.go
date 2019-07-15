package gotlsconfig

import (
	"crypto/tls"
	"time"
)

// New generates a new selfsinged rsa4096 certificate for subjectNames with 10 years expiry
func New(subjectNames ...string) *tls.Config {
	if len(subjectNames) == 0 {
		subjectNames = []string{"localhost"}
	}
	ssconfig := SelfSignedConfig{
		SAN:     subjectNames,
		KeyType: RSA4096,
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
