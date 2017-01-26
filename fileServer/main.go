package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
	"io/ioutil"
)

func main() {
	pubKeyBytes, _ := ioutil.ReadFile("public_key.pem")
	authServ := auth.Init()
	authServ.Register("file", pubKeyBytes)

	iris.Post("/get_file_hash", getFileHash)
	iris.Post("/write_file", writeFile)
	iris.Post("/read_file", readFile)
	iris.Post("/receive_goss", receiveGoss)
	iris.Post("/put_goss", putGoss)
	iris.ListenTLS(":8088", "./cert.pem", "./new_key.pem")
}
