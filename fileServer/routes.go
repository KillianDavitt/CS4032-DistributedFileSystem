package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
	"os"
)

func getFileHash(ctx *iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML(iris.StatusOK, "NOT AUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	contents, err := ioutil.ReadFile("files/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(contents)
	ctx.HTML(iris.StatusOK, string(hash[:]))
}

func writeFile(ctx *iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML(iris.StatusOK, "NOT AUTHORISED")
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
	ctx.HTML(iris.StatusOK, "ok")
}

func readFile(ctx *iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML(iris.StatusOK, "NOT AUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	file, err := os.Open("files/" + filename)
	if err != nil {
		log.Print(err)
		ctx.HTML(iris.StatusOK, "File not found")
		return
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}
	ctx.HTML(iris.StatusOK, string(contents))
}
