package main

import (
	"github.com/kataras/iris"
)

func addAction(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Not yet implemented")
}

func endTransaction(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Not yet implemented")
}

func beginTransaction(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Okay")
}
