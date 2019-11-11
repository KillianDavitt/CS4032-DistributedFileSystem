package rsaUtil

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
	pubKeyPem, err := ioutil.ReadFile("auth.pub.pem")
	if err != nil {
		log.Fatal(err)
	}
	pubKeyBytes, _ := pem.Decode(pubKeyPem)
	pubKeyInterface, err := x509.ParsePKIXPublicKey(pubKeyBytes.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	var pubKey rsa.PublicKey
	pubKey = *pubKeyInterface.(*rsa.PublicKey)
	return &pubKey
}

func GetPrivKey() *rsa.PrivateKey {
	privKeyPem, err := ioutil.ReadFile("auth.key.pem")
	if err != nil {
		log.Fatal(err)
	}
	privKeyBytes, _ := pem.Decode(privKeyPem)
	privKeyInterface, err := x509.ParsePKCS8PrivateKey(privKeyBytes.Bytes)
	if err != nil {
		log.Print("Error parsing priv key from file")
		log.Fatal(err)
	}
	var privKey *rsa.PrivateKey
	privKey = &*privKeyInterface.(*rsa.PrivateKey)

	return privKey

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

//VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
func Verify(pubKey *rsa.PublicKey, data []byte, signature []byte) bool {
	h := sha512.New()
	h.Write(data)
	hashed := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA512, hashed, signature)
	if err != nil {
		return false
	}
	return true
}
