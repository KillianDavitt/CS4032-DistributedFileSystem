package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Post("/lock_file", lockFile)
	iris.ListenTLS(":8086", "./cert.pem", "./key.pem")
}
