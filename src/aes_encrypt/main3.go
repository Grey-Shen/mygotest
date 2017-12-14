package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 128)
	if err != nil {
		log.Println(err)
		return
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Fatal(err)
		return
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&publicKeyBlock))

	fmt.Println(privateKeyPem)
	fmt.Println(publicKeyPem)

	f, err := os.OpenFile("rsa.yml", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("release:\n")
	f.WriteString("  portal_private_key: |\n")
	for _, v := range bytes.Split(pem.EncodeToMemory(&privateKeyBlock), []byte("\n")) {
		f.WriteString("                      ")
		if _, err := f.Write(v); err != nil {
			log.Fatal(err)
		} else {
			f.WriteString("\n")
		}

	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
