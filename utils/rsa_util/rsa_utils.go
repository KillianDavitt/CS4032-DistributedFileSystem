package rsa_util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func GetPubKey() *rsa.PublicKey {
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
	return &pubKey
}

func GetPrivKey() *rsa.PrivateKey {
	privKeyPem, err := ioutil.ReadFile("key_new.pem")
	if err != nil {

	}
	privKeyBytes, _ := pem.Decode(privKeyPem)
	privKeyInterface, err := x509.ParsePKCS1PrivateKey(privKeyBytes.Bytes)
	if err != nil {
		log.Print("Error parsing priv key from file")
		log.Fatal(err)
	}
	//var privKey *rsa.PrivateKey
	//privKey = &*privKeyInterface.(*rsa.PrivateKey)
	return privKeyInterface

}

//func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
func Sign(data []byte, priv *rsa.PrivateKey) []byte {
	h := sha512.New()
	h.Write(data)
	data_hash := h.Sum(nil)
	rng := rand.Reader
	signature, err := rsa.SignPKCS1v15(rng, priv, crypto.SHA512, data_hash[:])
	if err != nil {
		log.Print("Error signing")
		log.Fatal(err)
	}
	return signature
}
