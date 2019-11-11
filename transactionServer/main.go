package main

import (
	"io/ioutil"

	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
)

func main() {
	pubKeyBytes, _ := ioutil.ReadFile("public_key.pem")
	authServ := auth.Init()
	authServ.Register("transaction", pubKeyBytes)

	app := iris.New()

	app.Post("/put", put)
	app.Post("/end_transaction", endTransaction)
	app.Post("/begin_transaction", beginTransaction)

	app.Run(iris.TLS(":8080", "./transaction.crt.pem", "./transaction.key.pem"), iris.WithoutServerError(iris.ErrServerClosed))
}
