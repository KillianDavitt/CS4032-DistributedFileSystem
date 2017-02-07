package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris"
	"io/ioutil"

)

func main() {
	pubKeyBytes, _ := ioutil.ReadFile("public_key.pem")
	authServ := auth.Init()
	authServ.Register("transaction", pubKeyBytes)

	iris.Post("/put", put)
	
	iris.Post("/end_transaction", endTransaction)
	iris.Post("/begin_transaction", beginTransaction)
	iris.ListenTLS(":8080", "./transaction.crt.pem", "./transaction.key.pem")
}
