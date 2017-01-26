package main

import (
	"gopkg.in/redis.v5"
	"io/ioutil"
	"crypto/sha256"
)

func getFileRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 7})
}

// 
func isFileOutdated(filename string) bool {

}

func getFileHash(filename string) []byte {
	fileBytes, err := ioutil.ReadFile("files/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	hashBytes := sha256.Sum(fileBytes)
	return hashBytes[:]
}

func updateFilehistory(filename string, num int) {
	client := getFileRedis()
	err = client.Set(filename, string(num), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}
