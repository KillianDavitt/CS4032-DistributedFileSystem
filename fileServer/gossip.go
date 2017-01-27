package main

import (
	"encoding/json"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"net"
	"net/url"
	"strconv"
	"fmt"
)

// Receive a question asking if we need the goss
func receiveGoss(ctx *iris.Context) {
	filename := ctx.FormValue("filename")
	historyNumber := ctx.FormValue("history_number")
	historyInt, err := strconv.Atoi(historyNumber)
	if err != nil {
		log.Fatal(err)
	}
	if isFileOutdated(filename, historyInt) {
		go findGossRecipients(filename)
		ctx.HTML(iris.StatusOK, "yes")
	} else {
		ctx.HTML(iris.StatusOK, "no")
	}
}

// Accept the latest goss
func putGoss(ctx *iris.Context) {
	filename := ctx.FormValue("filename")
	fileBytes := []byte(ctx.FormValue("file"))
	err := ioutil.WriteFile(filename, fileBytes, 0777)
	if err != nil {
		log.Fatal(err)
	}
	ctx.HTML(iris.StatusOK, "Goss recv")
}

func getDirIp() net.IP {

	authServ := auth.Init()
	client := authServ.Client
	resp, err := client.Get("https://"+authServ.Ip.String()+":8080/get_dir_ip")
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBytes))
	dirIp := net.ParseIP(string(respBytes))
	fmt.Println(dirIp.String())
	return dirIp
}

func findGossRecipients(filename string) {
	authServ := auth.Init()
	client := authServ.Client
	dirIp := getDirIp()
	resp, err := client.PostForm("https://"+dirIp.String()+":8080/get_goss_servers", url.Values{"filename": {filename}})
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
	resp, err := client.PostForm("https://"+fileServerIp.String()+":8080/receive_goss", url.Values{"filename": {filename}})
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
		_, err := client.PostForm("https://"+fileServerIp.String()+":8080/put_goss", url.Values{"filename": {filename}})
		if err != nil {
			log.Fatal(err)
		}
	}
}
