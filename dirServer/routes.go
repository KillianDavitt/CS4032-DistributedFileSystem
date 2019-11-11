package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris/v12"
	"gopkg.in/redis.v5"
)

func registerFileHolder(ctx iris.Context) {
	// TODO
	if false {
		ctx.HTML("Forbidden")
		return
	}
	filename := ctx.FormValue("filename")
	fileHash := ctx.FormValue("file_hash")
	newFileServerIP := "0.0.0.0"
	NewRedisFile(filename, net.ParseIP(newFileServerIP), []byte(fileHash))
	ctx.HTML("Registered the file server as holding this file")
}

func getFile(ctx iris.Context) {
	if !isAllowed(ctx) {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.HTML("Invalid Token")
		return
	}

	filename := ctx.FormValue("filename")
	fileString := lookupFileName(filename)
	fileObj := UnmarshalFile([]byte(fileString))
	fmt.Println(fileObj.Ip.String())
	ctx.HTML(fileObj.Ip.String())
}

func putFile(ctx iris.Context) {
	if !isAllowed(ctx) {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.HTML("Invalid Token")
		return
	}
	fileName := ctx.FormValue("filename")
	fileHash := []byte(ctx.FormValue("hash"))
	fileJsonBytes := lookupFileName(fileName)
	if fileJsonBytes != "" {
		fmt.Println("Put: File existed")
		var fileObj file
		err := json.Unmarshal([]byte(fileJsonBytes), &fileObj)
		if err != nil {
			log.Fatal(err)
		}

		fileObj.Hash = fileHash
		fileObj.UpdateRedisFile()

		ctx.HTML(fileObj.Ip.String())
	} else {
		// TODO: Fix this
		fmt.Println("Put: File did not exist")
		newFileServerIP := "0.0.0.0"
		NewRedisFile(fileName, net.ParseIP(newFileServerIP), fileHash)
		ctx.HTML(newFileServerIP)
	}
}

func getFileRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 3})
}

func lookupFileName(filename string) string {
	fileClient := getFileRedis()
	res, _ := fileClient.Get(filename).Result()
	return res
}

func listFiles(ctx iris.Context) {
	if !isAllowed(ctx) {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.HTML("Invalid token")
		return
	}
	fileClient := getFileRedis()
	keys, err := fileClient.Keys("*").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(keys)
	jsonFiles, err := json.Marshal(keys)
	fmt.Println(string(jsonFiles))
	if err != nil {
		log.Fatal(err)
	}
	ctx.HTML(string(jsonFiles))
}

func registerToken(ctx iris.Context) {
	pubKey := auth.RetrieveKey("authserver")
	token := ctx.FormValue("token")
	ticket := ticket.GetTicketMap(token, pubKey)
	token_client := getTokenRedis()
	expiryString, err := ticket.Expiry_date.MarshalText()
	if err != nil {
		log.Fatal(err)
	}
	err = token_client.Set(string(ticket.Token), string(expiryString), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	ctx.HTML("Register token succ")
}

func isAllowed(ctx iris.Context) bool {
	token := ctx.FormValue("token")
	pubKey := auth.RetrieveKey("authserver")
	ticket := ticket.GetTicketMap(token, pubKey)
	client := getTokenRedis()
	_, err := client.Get(string(ticket.Token)).Result()
	if err != nil {
		return false
	}
	return true
}

func getTokenRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 1})
}
