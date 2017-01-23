package main

import (
	"github.com/kataras/iris"
)

func beginTransaction(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Okay")
}
