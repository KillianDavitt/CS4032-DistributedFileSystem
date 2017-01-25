package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

func get(args []string, client *http.Client, ip net.IP, ticketMapBytes []byte) {
	// Contact the dir server and get the ip of a file server
	filename := args[0]
	fmt.Println(filename)
	// TODO Caching would go here
	

	resp, err := client.PostForm("https://"+ip.String()+":8089/get_file", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	fileserverIp := net.ParseIP(string(respBytes))

	if isFileCached(filename, client, fileserverIp, ticketMapBytes) {
		fmt.Println(string(getCachedFile(filename, client, fileserverIp, ticketMapBytes)))
		return
	}


	fmt.Println(string(respBytes))
	// Now put to the file server on the ip we received
	resp, err = client.PostForm("https://"+fileserverIp.String()+":8088/read_file", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBytes))
}
