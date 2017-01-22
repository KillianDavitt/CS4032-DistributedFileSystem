package main

import (
	"github.com/kataras/iris"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
)

func main() {
	// First thing, get contact with the auth server and confirm it's indentity
	auth.Init()
	iris.Get("/get_file", getFile)
	iris.Post("/list_files", listFiles)
	iris.Post("/put_file", putFile)
	iris.Post("/register_token", registerToken)
	iris.ListenTLS(":8089", "./cert.pem", "./new_key.pem")
}
