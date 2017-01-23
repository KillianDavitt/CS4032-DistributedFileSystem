package main

import (
	"github.com/kataras/iris"
)

func main() {
	auth()
	iris.Post("/begin_transaction", beginTransaction)
	iris.ListenTLS(":8088", "./cert.pem", "./new_key.pem")
}
