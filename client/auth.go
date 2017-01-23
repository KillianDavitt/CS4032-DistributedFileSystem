package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func auth(authServ *authServer) {
	// InsecureSkipVerify must be set since we need to contact the auth server once to find it's fingerprint
	conn, err := tls.Dial("tcp", authServ.Ip.String()+":8080", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	CA_Pool := x509.NewCertPool()

	fmt.Println("About to print sig")
	servedCert := *conn.ConnectionState().PeerCertificates[0]
	servedPubKey := rsa.PublicKey(*servedCert.PublicKey.(*rsa.PublicKey))

	fmt.Println(authServ.PubKey.E)
	fmt.Println(servedPubKey.E)

	if authServ.PubKey == servedPubKey {
		fmt.Println("Keys match")
	} else {
		fmt.Println("You have not saved this auth servers public key...")
		fmt.Println(servedPubKey.E)
		fmt.Println("Would you like to accept this key?")
		input := ""
		fmt.Scanf("%s", &input)
		if input == "y" {
			authServ.PubKey = servedPubKey
			writeConfig(authServ)
			fmt.Println("This auth server public key has been accepted")
		}
	}
	CA_Pool.AddCert(&servedCert)

	// This bit set to insecure until I can fix it later
	tlsConf := &tls.Config{RootCAs: CA_Pool, InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: transport}

	username := ""
	password := ""
	fmt.Println("Enter your username")
	fmt.Scanf("%s", &username)
	fmt.Println("enter your password")
	fmt.Scanf("%s", &password)
	resp, err := client.PostForm("https://"+authServ.Ip.String()+":8080/login", url.Values{"username": {username}, "password": {password}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ticket := ticket.GetTicketMap(string(bytes), &authServ.PubKey)
	fmt.Println(ticket)

}
