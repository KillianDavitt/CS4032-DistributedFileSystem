package main

import (
	"github.com/kataras/iris"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"log"
	"net"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

// Receive a question asking if we need the goss
func receiveGoss(ctx *iris.Context) {
	filename := ctx.FormValue("filename")
	go findGossRecipients(filename)
	ctx.HTML(iris.StatusOK, "Goss sent")
}

// Accept the latest goss
func putGoss(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Goss recv")
}

func getDirIp() net.IP {
	client := auth.GetTLSClient()
	authServ := auth.Init()
	resp, err := client.PostForm("https://" + authServ.Ip.String() + ":8080/get_dir_ip", url.Values{})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	dirIp := net.ParseIP(string(respBytes))
	return dirIp
}

func findGossRecipients(filename string) {
	client := auth.GetTLSClient()
	dirIp := getDirIp()
	resp, err := client.PostForm("https://" + dirIp.String() + ":8080/get_goss_servers", url.Values{"filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	var fileIps []net.IP
	err = json.Unmarshal(respBytes, fileIps)

	for _, ip := range fileIps {
		go sendGoss(ip, filename)
	}
}

func sendGoss(fileServerIp net.IP, filename string) {
	client := auth.GetTLSClient()
	// Ask if they need it,
	resp, err := client.PostForm("https://" + fileServerIp.String() + ":8080/receive_goss", url.Values{"filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	respString := string(respBytes)
	if respString != "yes" {
		// fileServer does not want the latest goss
		return
	} else {
		// Send the latest goss
		_, err := client.PostForm("https://" + fileServerIp.String() + ":8080/put_goss", url.Values{"filename": {filename}})
		if err != nil {
			log.Fatal(err)
		}
	}
}

