package main

// Sadkey is a command line tool for generating RSA public and
// private keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {

	arg := flag.String("keygen", "", "sadkey -keygen rsa")
	flag.Parse()

	args := os.Args[:]

	if len(args) <= 1 {
		flag.Usage()
		return
	}

	if *arg == "rsa" {
		generateRSAKeyPairs()
		fmt.Println("keys generated")
		return
	} else {
		flag.Usage()
	}
}

func generateRSAKeyPairs() {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	privPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priv),
	}

	f, err := os.Create("privPemFile.pem")
	if err != nil {
		fmt.Println(err)
		return
	}
	pr := pem.EncodeToMemory(privPemBlock)
	f.Write(pr)
	f.Close()

	pub := &priv.PublicKey

	pubPemBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(pub),
	}

	nf, err := os.Create("pubPemFile.pem")
	if err != nil {
		fmt.Println(err)
		return
	}

	ppr := pem.EncodeToMemory(pubPemBlock)
	nf.Write(ppr)
	nf.Close()

}
