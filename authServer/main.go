package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Post("/login", login)

	iris.ListenTLS(":8080", "./root-ca_mod.crt", "./root-key_mod.pem")
}
