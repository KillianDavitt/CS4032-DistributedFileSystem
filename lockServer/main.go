package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
)

func main() {
	auth.Init()
	iris.Post("/lock_file", lockFile)
	iris.Post("/unlock_file", unlockFile)
	iris.ListenTLS(":8086", "./cert.pem", "./key.pem")
}
