package auth

import (
	"log"
	"net/url"
)

func (serv *AuthServer) Register(serverType string, pubKeyBytes []byte) {
	_, err := serv.Client.PostForm("https://"+serv.Ip.String()+":8080/register_server", url.Values{"server_type": {serverType}, "public_key": {string(pubKeyBytes)}})
	if err != nil {
		log.Fatal(err)
	}
}
