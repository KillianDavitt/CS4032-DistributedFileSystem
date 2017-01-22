package main

import (
	"fmt"
	"net/http"
	"net"
	"net/url"
	"log"
	"io/ioutil"
	"encoding/json"
	//"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
)

func list(_ []string, client *http.Client, ip net.IP, ticketMapBytes []byte){
	//ticketMapString := myTicket.MarshalTicket()
	resp, err := client.PostForm("https://" + "0.0.0.0" + ":8089/list_files", url.Values{"token": {string(ticketMapBytes)}})
	
	if err != nil {
		log.Fatal(err)
	}
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var fileList []string
	err = json.Unmarshal(responseBytes, &fileList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileList)
}
