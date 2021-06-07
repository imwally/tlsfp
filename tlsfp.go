package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
)

func errAndExit(err error) {
	fmt.Fprintf(os.Stderr, "tlsfp: %s\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		errAndExit(errors.New("expected at least 1 argument"))
	}

	host := os.Args[1]
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		errAndExit(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	cert := state.PeerCertificates[0]
	fmt.Printf("SHA1:\t% X\n", sha1.Sum(cert.Raw))
	fmt.Printf("SHA256:\t% X\n", sha256.Sum256(cert.Raw))
}
