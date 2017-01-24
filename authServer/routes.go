package main

import (
	"crypto/tls"
	"crypto/rsa"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsa_util"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
	"log"
	"net/http"
	"net/url"
	"net"
	"fmt"
	"encoding/pem"
	"crypto/x509"
)

func getDirIp(ctx *iris.Context) {
	dirServerIps := getDirIps()
	dirServerIp := dirServerIps[0]
	ctx.HTML(iris.StatusOK, string(dirServerIp))
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
	serverTypeString := ctx.FormValue("server_type")
	serverType := FILE
	if serverTypeString == "fileserver" {
		serverType = FILE
	} else {
		serverType = DIR
	}
	
	pubKeyPem := ctx.FormValue("public_key")

	block, _ := pem.Decode([]byte(pubKeyPem))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pubKeyInter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	pubKey := pubKeyInter.(*rsa.PublicKey)
	
	
	serverIP := net.ParseIP(ctx.Request.RemoteAddr)
	fmt.Println("A server wants to register itself with the following public key\n")
	fmt.Println(pubKey)
	fmt.Println("\nWould you like to accept? (y/n)")
	inp := ""
	fmt.Scanf("%s", &inp)
	if inp == "y" {
		serv := NewServer(serverIP, serverType, pubKey)
		serv.writeServerRedis()
	}
	ctx.HTML(iris.StatusOK, "Not registered")
}
