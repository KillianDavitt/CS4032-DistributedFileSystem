package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
)

func getFileHash(ctx iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML("NOT AUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	contents, err := ioutil.ReadFile("files/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(contents)
	ctx.HTML(string(hash[:]))
}

func writeFile(ctx iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML("NOT AUTHORISED")
		return
	}
	// this needs to be run before a new file is put
	fileString := ctx.FormValue("file")
	fileBytes := []byte(fileString)
	filename := ctx.FormValue("filename")
	file, err := os.Create("files/" + filename)
	if err != nil {
		file, err = os.Create("files/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(err)
	}
	defer file.Close()
	fmt.Println(fileBytes)
	_, err = file.Write(fileBytes)
	if err != nil {
		log.Fatal(err)
	}
	go findGossRecipients(filename)
	ctx.HTML("ok")
}

func readFile(ctx iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML("NOT AUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	file, err := os.Open("files/" + filename)
	if err != nil {
		log.Print(err)
		ctx.HTML("File not found")
		return
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}
	ctx.HTML(string(contents))
}
