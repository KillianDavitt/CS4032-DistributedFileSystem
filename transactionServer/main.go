package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
)

func main() {
	auth.Init()
	iris.Post("/begin_transaction", beginTransaction)
	iris.ListenTLS(":8088", "./cert.pem", "./new_key.pem")
}
