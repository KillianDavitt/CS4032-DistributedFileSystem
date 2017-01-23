package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Post("/write_file", writeFile)
	iris.Post("/read_file", readFile)
	iris.ListenTLS(":8088", "./cert.pem", "./new_key.pem")
}
