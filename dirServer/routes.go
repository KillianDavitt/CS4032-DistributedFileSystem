package main
import (
	"github.com/kataras/iris"
	"log"
	"io/ioutil"
)
func getFile(ctx *iris.Context){
	filename := ctx.URLParam("filename")
	f := readFiles()
	ip := f.getFile(filename).String()
	ctx.HTML(iris.StatusOK, ip)
}

func putFile(ctx *iris.Context){

}


func listFiles(ctx *iris.Context){
	f := readFiles()
	ctx.HTML(iris.StatusOK, f.getFile("test.txt").String())
}

func registerToken(ctx *iris.Context){

}
