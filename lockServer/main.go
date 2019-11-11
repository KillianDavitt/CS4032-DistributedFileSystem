package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
)

func main() {
	auth.Init()

	app := iris.New()

	app.Post("/lock_file", lockFile)
	app.Post("/unlock_file", unlockFile)

	app.Run(iris.TLS(":8080", "./cert.pem", "./key.pem"), iris.WithoutServerError(iris.ErrServerClosed))
}
