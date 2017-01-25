package main

import (
	"log"
	"os"
	"io/ioutil"
	"crypto/sha256"
	"net/http"
	"net"
	"net/url"
)

func initCache() {
	os.RemoveAll("cache/")
	os.MkdirAll("cache/", os.ModeDir)
}

func isFileCached(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) (bool) {
	if _, err := os.Stat("cache/" + filename); os.IsNotExist(err) {
		return false
	}

	if string(getLocalFileHash(filename)) == string(getRemoteFileHash(filename, client, fileServerIp, ticketMapBytes)) {
		return true
	} else {
		return false
	}
}

func getCachedFile(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) ([]byte) {
	if !isFileCached(filename, client, fileServerIp, ticketMapBytes) {
		log.Fatal("Requested file isnt cached")
	}
	fileBytes, err := ioutil.ReadFile("cache/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	return fileBytes
}

func removeOldFile(filename string) {
	os.Remove("cache/" + filename)
}

func getLocalFileHash(filename string) ([]byte) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(fileBytes)
	return hash[:]
}

func getRemoteFileHash(filename string, client *http.Client, fileServerIp net.IP, ticketMapBytes []byte) ([]byte) {
	resp, err := client.PostForm("https://" + fileServerIp.String() + ":8089/get_file_hash", url.Values{"token": {string(ticketMapBytes)}, "filename": {filename}})
	if err != nil {
		log.Fatal(err)
	}
	var hashBytes []byte
	_, err = resp.Body.Read(hashBytes)
	return hashBytes[:]
}
