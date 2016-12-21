package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Post("/write_file", writeFile)
	iris.Get("/read_file", readFile)
	iris.ListenTLS(":8080", "./cert.pem", "./key.pem")
}
