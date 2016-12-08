package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/write_file", writeFile)

	iris.ListenTLS(":8080", "./cert.pem", "./key.pem")
}
