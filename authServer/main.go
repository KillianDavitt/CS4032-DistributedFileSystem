package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
	"net/http"
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

	// Gen token, give back to user, then give to all servers
	ticket := newTicket()
	out, err := json.Marshal(ticket)
	fmt.Printf("%#v", ticket)
	fmt.Println(string(ticket.token))
	c.HTML(iris.StatusOK, fmt.Sprintf("%#v", ticket))
	dirServerIP := "10.1.2.1"
	_, err = http.Get("https://" + dirServerIP + "/register_token" + "?token=" + string(out))
}
