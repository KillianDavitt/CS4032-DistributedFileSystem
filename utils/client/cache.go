package client

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

func initCache() {
	os.RemoveAll("cache/")
	os.MkdirAll("cache/", 1777)
}

func isFileCached(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) bool {
	if _, err := os.Stat("cache/" + filename); os.IsNotExist(err) {
		return false
	}
	fmt.Println(getLocalFileHash(filename))
	fmt.Println(getRemoteFileHash(filename, client, fileServerIp, ticketMapBytes))
	if string(getLocalFileHash(filename)) == string(getRemoteFileHash(filename, client, fileServerIp, ticketMapBytes)) {
		return true
	} else {
		return false
	}
}

func getCachedFile(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) []byte {
	if !isFileCached(filename, client, fileServerIp, ticketMapBytes) {
		log.Fatal("Requested file isnt cached")
	}
	fileBytes, err := ioutil.ReadFile("cache/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded cached file....")
	return fileBytes
}

func removeOldFile(filename string) {
	os.Remove("cache/" + filename)
}

func getLocalFileHash(filename string) string {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(fileBytes)
	return string(hash[:])
}

func getRemoteFileHash(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) string {
	resp, err := client.PostForm("https://"+fileServerIp.String()+":8088/get_file_hash", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}

	hashBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashBytes[:])
}
