package gotlsconfig

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"
)

// Constants for all key types we support.
const (
	EC256   = KeyType("P256")
	EC384   = KeyType("P384")
	RSA2048 = KeyType("2048")
	RSA4096 = KeyType("4096")
	RSA8192 = KeyType("8192")
)

type KeyType string

// SelfSignedConfig configures a self-signed certificate.
type SelfSignedConfig struct {
	SAN          []string // Subject Alternative Names
	KeyType      KeyType
	Expire       time.Time
	Organization string
}

// newSelfSignedCertificate returns a new self-signed certificate.
func newSelfSignedCertificate(ssconfig SelfSignedConfig) (tls.Certificate, error) {
	// start by generating private key
	var privKey interface{}
	var err error
	switch ssconfig.KeyType {
	case "", EC256:
		privKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case EC384:
		privKey, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case RSA2048:
		privKey, err = rsa.GenerateKey(rand.Reader, 2048)
	case RSA4096:
		privKey, err = rsa.GenerateKey(rand.Reader, 4096)
	case RSA8192:
		privKey, err = rsa.GenerateKey(rand.Reader, 8192)
	default:
		return tls.Certificate{}, fmt.Errorf("cannot generate private key; unknown key type %v", ssconfig.KeyType)
	}
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to generate private key: %v", err)
	}

	// create certificate structure with proper values
	notBefore := time.Now()
	notAfter := ssconfig.Expire
	if notAfter.IsZero() || notAfter.Before(notBefore) {
		notAfter = notBefore.Add(24 * time.Hour * 7)
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("failed to generate serial number: %v", err)
	}
	if ssconfig.Organization == "" {
		ssconfig.Organization = "github.com/mschneider82/gotlsconfig self-signed"
	}
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      pkix.Name{Organization: []string{ssconfig.Organization}},
		NotBefore:    notBefore,
		NotAfter:     notAfter,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	if len(ssconfig.SAN) == 0 {
		ssconfig.SAN = []string{""}
	}
	for _, san := range ssconfig.SAN {
		if ip := net.ParseIP(san); ip != nil {
			cert.IPAddresses = append(cert.IPAddresses, ip)
		} else {
			cert.DNSNames = append(cert.DNSNames, strings.ToLower(san))
		}
	}

	// generate the associated public key
	publicKey := func(privKey interface{}) interface{} {
		switch k := privKey.(type) {
		case *rsa.PrivateKey:
			return &k.PublicKey
		case *ecdsa.PrivateKey:
			return &k.PublicKey
		default:
			return fmt.Errorf("unknown key type")
		}
	}
	derBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, publicKey(privKey), privKey)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("could not create certificate: %v", err)
	}

	chain := [][]byte{derBytes}

	return tls.Certificate{
		Certificate: chain,
		PrivateKey:  privKey,
		Leaf:        cert,
	}, nil
}
