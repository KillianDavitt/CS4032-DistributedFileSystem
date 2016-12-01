package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func getPubKey() rsa.PublicKey {
	pubKeyPem, err := ioutil.ReadFile("pubkey.pem")
	pubKeyBytes, _ := pem.Decode(pubKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	pubKeyInterface, err := x509.ParsePKIXPublicKey(pubKeyBytes.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	var pubKey rsa.PublicKey
	pubKey = *pubKeyInterface.(*rsa.PublicKey)
	return pubKey
}
