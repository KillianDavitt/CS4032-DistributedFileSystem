package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	app.Post("/login", login)
	app.Post("/register_server", registerServer)
	app.Get("/get_dir_ip", getDirIp)

	app.Run(iris.TLS(":8080", "./auth.crt.pem", "./auth.key.pem"), iris.WithoutServerError(iris.ErrServerClosed))
}
