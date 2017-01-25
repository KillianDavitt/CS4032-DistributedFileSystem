package main

import (
	"github.com/kataras/iris"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
)

func main() {
	auth.Init()
	iris.Post("/lock_file", lockFile)
	iris.Post("/unlock_file", unlockFile)
	iris.ListenTLS(":8086", "./cert.pem", "./key.pem")
}
