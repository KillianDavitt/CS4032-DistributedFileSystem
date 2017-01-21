package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsa_util"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
	"net/http"
	"fmt"
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
	
	privKey := rsa_util.GetPrivKey()

	dirServerIP := "10.1.2.1"


	ticketMapString := new_ticket.CreateTicketMap(privKey)
	c.HTML(iris.StatusOK, ticketMapString)
	// Send token to the dir server
	
	go func() {
		_, err = http.Get("https://" + dirServerIP + "/register_token" + "?token=" + ticketMapString)
	}()
}
