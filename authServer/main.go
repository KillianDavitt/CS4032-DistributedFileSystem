package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/login", login)

	iris.ListenTLS(":8080", "./cert.pem", "./key_new.pem")
}
