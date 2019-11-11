package main

import (
	"io/ioutil"
	"log"

	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
)

func main() {
	pubKeyBytes, err := ioutil.ReadFile("./fs.pub.pem")
	if err != nil {
		log.Fatal(err)
	}

	authServ := auth.Init()
	authServ.Register("file", pubKeyBytes)

	app := iris.New()

	app.Post("/get_file_hash", getFileHash)
	app.Post("/write_file", writeFile)
	app.Post("/read_file", readFile)
	app.Post("/receive_goss", receiveGoss)
	app.Post("/put_goss", putGoss)

	app.Run(iris.TLS(":8080", "./fs.cert.pem", "./fs.key.pem"), iris.WithoutServerError(iris.ErrServerClosed))
}
