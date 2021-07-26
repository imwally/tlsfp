package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"os"
)

func errAndExit(err error) {
	if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
		err = wrappedErr
	}

	fmt.Fprintf(os.Stderr, "tlsfp: %s\n", err)
	os.Exit(1)
}

func getCert(addr string) *x509.Certificate {
	conn, err := tls.Dial("tcp", addr+":443", nil)
	if err != nil {
		errAndExit(err)
	}
	defer conn.Close()

	state := conn.ConnectionState()

	return state.PeerCertificates[0]
}

func main() {
	var algo int

	tlsfs := flag.NewFlagSet("tlsfp", flag.ExitOnError)
	tlsfs.IntVar(&algo, "a", 1, "algorithm: 1, 256")
	tlsfs.Parse(os.Args[1:])

	if len(tlsfs.Args()) < 1 {
		errAndExit(errors.New("no host specified"))
	}

	switch algo {
	case 1:
		cert := getCert(tlsfs.Arg(0))
		fmt.Printf("% X\n", sha1.Sum(cert.Raw))
	case 256:
		cert := getCert(tlsfs.Arg(0))
		fmt.Printf("% X\n", sha256.Sum256(cert.Raw))
	default:
		errText := fmt.Sprintf("unrecognized algorithm: %d", algo)
		errAndExit(errors.New(errText))
	}
}
