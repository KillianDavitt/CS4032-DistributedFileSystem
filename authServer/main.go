package main

import (
	"github.com/kataras/iris"
)

func main() {
	
	iris.Post("/login", login)
	iris.Post("/register_server", registerServer)
	iris.Get("/get_dir_ip", getDirIp)
	iris.ListenTLS(":8080", "./auth.crt.pem", "./auth.key.pem")
}
