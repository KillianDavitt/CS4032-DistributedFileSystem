package main

import (
	"log"
	"crypto/tls"
	"fmt"
)

func main(){
	conn, err := tls.Dial("tcp", "mail.google.com:443", &tls.Config{})
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%x", conn.ConnectionState().PeerCertificates[0].Signature)
}
