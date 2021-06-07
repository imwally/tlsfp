package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "expected at lest 1 argument\n")
		os.Exit(-1)
	}

	host := os.Args[1]
	conn, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()
	cert := state.PeerCertificates[0]
	fmt.Printf("SHA1:\t% X\n", sha1.Sum(cert.Raw))
	fmt.Printf("SHA256:\t% X\n", sha256.Sum256(cert.Raw))
}
