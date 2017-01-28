package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
	"io/ioutil"
	"log"
)

func main() {

	pubKeyBytes, err := ioutil.ReadFile("dirServer/dir.pub.pem")
	if err != nil {
		log.Fatal(err)
	}
	
	// Init contacts the auth server and organises OUR trust of it
	authServer := auth.Init()
	// Register contacts the auth server and organises THEIR trust of us
	// It also confirms to the auth server that we are currently online and acting as a dir server
	authServer.Register("dirServer", pubKeyBytes)

	// Our own routes
	iris.Post("/get_file", getFile)
	iris.Post("/list_files", listFiles)
	iris.Post("/put_file", putFile)
	iris.Post("/register_token", registerToken)
	iris.ListenTLS(":8089", "dirServer/dir.crt.pem", "dirServer/dir.key.pem")
}
