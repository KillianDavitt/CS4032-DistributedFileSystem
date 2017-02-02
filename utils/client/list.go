package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	//"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
)

func List(_ []string, client *http.Client, ip net.IP, ticketMapBytes []byte) {
	//ticketMapString := myTicket.MarshalTicket()
	resp, err := client.PostForm("https://"+ip.String()+":8089/list_files", url.Values{"token": {string(ticketMapBytes)}})

	if err != nil {
		fmt.Println("Can't establish connection to the directory server")
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
