package main

import (
	"crypto/tls"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsa_util"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
	"net/url"
)

func getDirIp(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "0.0.0.0")
}

func getLoginRedis() (*redis.Client) {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB:  0, })
}

func login(c *iris.Context) {
	// Connect to redis
	client := getLoginRedis()

	username := c.FormValue("username")
	password := c.FormValue("password")

	hashedPassword, err := client.Get(username).Result()

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.HTML(iris.StatusForbidden, "Incorrect username or password")
	}

	// Gen token, give back to user, then give to all servers
	new_ticket := ticket.NewTicket()
	privKey := rsa_util.GetPrivKey()
	ticketMapString := new_ticket.CreateTicketMap(privKey)

	distributeTickets(ticketMapString)
	c.HTML(iris.StatusOK, ticketMapString)
}

func distributeTickets(ticketMapString string) {
	serverIps := getServerIps()
	tlsConf := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: transport}
	for _, ip := range serverIps {
		_, err := client.PostForm("https://"+ ip.String() + ":8089/register_token", url.Values{"token": {ticketMapString}})
		if err != nil {
			log.Fatal(err)
		}

	}
}

func registerServer(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Ok")
}
