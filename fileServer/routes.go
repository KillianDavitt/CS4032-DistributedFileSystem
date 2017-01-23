package main

import (
	"github.com/kataras/iris"
	"log"
	"os"
	"fmt"
	"io/ioutil"
)

func writeFile(ctx *iris.Context){
	// this needs to be run before a new file is put
	fileString := ctx.FormValue("file")
	fileBytes := []byte(fileString)
	filename := ctx.FormValue("filename")
	file, err := os.Create(filename)
	if err != nil {
		file, err = os.Create(filename)
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
	ctx.HTML(iris.StatusOK, "ok")
}

func readFile(ctx *iris.Context){

	filename := ctx.FormValue("filename")
	file, err := os.Open(filename)
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
