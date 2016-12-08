package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsa_util"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
)

func writeFile(ctx *iris.Context){
	// this needs to be run before a new file is put
	info, err := ctx.FormFile("file")
	if err != nil {
		log.Print(err)
	}
	file, err := info.Open()
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	
	fileContents, err := ioutil.ReadAll(file)
	
	ctx.HTML(iris.StatusOK, string(fileContents))
}

func readFile(c *iris.Context){

}
