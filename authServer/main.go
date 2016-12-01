package main

import (
	"crypto/rand"
	"fmt"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	iris.Get("/login", login)

	iris.ListenTLS(":8080", "./cert.pem", "./key_new.pem")
}

func get_token() {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func login(c *iris.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	username := c.FormValueString("username")
	password := c.FormValueString("password")
	hashedPassword, err := client.Get(username).Result()
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.HTML(iris.StatusForbidden, "Incorrect username or password")
	}
}
