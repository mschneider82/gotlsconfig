package gotlsconfig_test

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/mschneider82/gotlsconfig"
)

func Example() {
	config := gotlsconfig.New("localhost")

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		connp, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conn := tls.Server(connp, config)
		go func(c net.Conn) {
			io.Copy(os.Stdout, c)
			fmt.Println()
			c.Close()
		}(conn)
	}
}
