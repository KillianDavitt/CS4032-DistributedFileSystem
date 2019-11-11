package main

import (
	"io/ioutil"
	"log"

	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
)

func main() {

	pubKeyBytes, err := ioutil.ReadFile("dir.pub.pem")
	if err != nil {
		log.Fatal(err)
	}

	// Init contacts the auth server and organises OUR trust of it
	authServer := auth.Init()
	// Register contacts the auth server and organises THEIR trust of us
	// It also confirms to the auth server that we are currently online and acting as a dir server
	authServer.Register("dirServer", pubKeyBytes)

	// Our own routes
	app := iris.New()

	app.Post("/get_file", getFile)
	app.Post("/list_files", listFiles)
	app.Post("/put_file", putFile)
	app.Post("/register_token", registerToken)

	app.Run(iris.TLS(":8080", "dir.crt.pem", "dir.key.pem"), iris.WithoutServerError(iris.ErrServerClosed))
}
