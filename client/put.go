package main

import (
	"net"
	"log"
	"crypto/sha256"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
)

func put(args []string, client *http.Client, ip net.IP, ticketMapBytes []byte){
	filename := args[0]
	fmt.Println(filename)
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)    
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	hash := hasher.Sum(nil)
	resp, err := client.PostForm("https://" + ip.String() + ":8089/put_file", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}, "hash": {string(hash)}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBytes))
}

