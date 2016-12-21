package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/login", login)

	iris.ListenTLS(":8080", "./newcert.pem", "./newkey_new.pem")
}
