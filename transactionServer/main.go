package main

import (
	"github.com/kataras/iris"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
)

func main() {
	auth.Init()
	iris.Post("/begin_transaction", beginTransaction)
	iris.ListenTLS(":8088", "./cert.pem", "./new_key.pem")
}
