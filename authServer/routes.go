package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsa_util"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
	"fmt"
	"encoding/base64"
)

func login(c *iris.Context) {

	// Connect to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println(username)
	hashedPassword, err := client.Get(username).Result()

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	// Invalid username or password, RIP
	if err != nil {
		c.HTML(iris.StatusForbidden, "Incorrect username or password")
	}

	// Gen token, give back to user, then give to all servers
	new_ticket := ticket.NewTicket()
	ticket_bson, err := bson.Marshal(&new_ticket)
	if err != nil {
		log.Fatal(err)
	}
	privKey := rsa_util.GetPrivKey()
	signed_ticket := rsa_util.Sign(ticket_bson, privKey)
	signed_ticket_64 := base64.StdEncoding.EncodeToString(signed_ticket)
	ticket_bson_64 := base64.StdEncoding.EncodeToString(ticket_bson)
	c.HTML(iris.StatusOK, signed_ticket_64 + "|" + ticket_bson_64)
	dirServerIP := "10.1.2.1"

	// Send token to the dir server
	go func() {
		_, err = http.Get("https://" + dirServerIP + "/register_token" + "?token=" + string(ticket_bson) + string(signed_ticket))
	}()
}
