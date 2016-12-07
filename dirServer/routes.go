package main
import (
	"fmt"
	"github.com/kataras/iris"
)
func getFile(ctx *iris.Context, f *files){
	filename := ctx.URLParam("filename")
	f := readFiles()
	ip := f.getFile(filename).String()
	ctx.HTML(iris.StatusOK, ip)
}

func putFile(ctx *iris.Context, f *files){
	// this needs to be run before a new file is put
	filename := ctx.URLParam("filename")
	f := readFiles()
}

func listFiles(ctx *iris.Context, f *files){
	f := readFiles()
	ctx.HTML(iris.StatusOK, f.getFile("test.txt").String())
}

func registerToken(ctx *iris.Context){

}
