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
)

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
	new_ticket := ticket.NewTicket()
	ticket_bson, err := bson.Marshal(&new_ticket)
	if err != nil {
		log.Fatal(err)
	}
	privKey := rsa_util.GetPrivKey()
	signed_ticket := rsa_util.Sign(ticket_bson, privKey)

	c.HTML(iris.StatusOK, string(signed_ticket)+string(ticket_bson))
	dirServerIP := "10.1.2.1"
	go func() {
		_, err = http.Get("https://" + dirServerIP + "/register_token" + "?token=" + string(ticket_bson) + string(signed_ticket))
	}()
}
