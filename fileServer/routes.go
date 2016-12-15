package main

import (
	"github.com/kataras/iris"
	"log"
	"os"
	"io/ioutil"
)

func writeFile(ctx *iris.Context){
	// this needs to be run before a new file is put
	info, err := ctx.FormFile("file")
	if err != nil {
		log.Print(err)
	}

	filename := info.Filename
	file, err := info.Open()
	if err != nil {
		log.Print(err)
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	
	newFile, err := os.Create("/home/killian/" + filename)
	if err != nil{
		log.Print(err)
	}
	_, err = newFile.Write(fileContents)
	if err != nil {
		log.Print(err)
	}
	ctx.HTML(iris.StatusOK, string(fileContents))
	newFile.Close()
}

func readFile(ctx *iris.Context){

	filename := ctx.URLParam("filename")
	log.Print(filename)
	file, err := os.Open("/home/killian/" + filename)
	if err != nil {
		log.Print(err)
		ctx.HTML(iris.StatusOK, "File not found")
	}

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print(err)
	}
	ctx.HTML(iris.StatusOK, string(contents))
	
}
