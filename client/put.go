package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

func put(args []string, client *http.Client, ip net.IP, ticketMapBytes []byte) {
	// Contact the dir server and get the ip of a file server
	filename := args[0]
	fmt.Println(filename)
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filename)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	hash := hasher.Sum(nil)
	resp, err := client.PostForm("https://"+ip.String()+":8089/put_file", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}, "hash": {string(hash)}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fileserverIp := net.ParseIP(string(respBytes))

	// Now put to the file server on the ip we received
	resp, err = client.PostForm("https://"+fileserverIp.String()+":8088/write_file", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}, "file": {string(s)}})
}
