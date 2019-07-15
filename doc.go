// Package gotlsconfig makes it easy to get a secure *tls.Config for testing and
// development. It generates on the fly pub/private certificates.
// This mitigates the usage of static private keys in go code.
//
// Its better to use self singed certificates instead of doing plain traffic!
package gotlsconfig
