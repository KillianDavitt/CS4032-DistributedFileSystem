package main

import (
	"fmt"
	"net/http"
	"net"
	"log"
)

func list(_ string, client *http.Client, ip net.IP){
	resp, err := client.Get("https://" + ip.String() + ":8089/list_files")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
