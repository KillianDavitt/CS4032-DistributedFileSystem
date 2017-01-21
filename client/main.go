package main

import (
	"net"
	"crypto/rsa"
	"fmt"
	"os"
	"gopkg.in/mgo.v2/bson"
	"log"
	"io/ioutil"
)

func help(){
	fmt.Println("Help:\nls\nput\nget")
}

func list(_ string){
	fmt.Println("Listing")
}

func put_file(_ string){
	fmt.Println("Putting file")
}

func transaction_start(_ string){
	fmt.Println("Starting transaction..")
}

func transaction_end(_ string){
	fmt.Println("End transaction")
}


type authServer struct{
	Ip net.IP
	PubKey rsa.PublicKey
}

func writeConfig(authServ *authServer) {
	fmt.Println(authServ.PubKey.E)
	authServBytes, err := bson.Marshal(authServ)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(".dfs.conf")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(authServBytes)
}

func getConfig() (*authServer) {
	if _, err := os.Stat(".dfs.conf"); os.IsNotExist(err) {
		newServ := &authServer{}
		fmt.Println("Enter the ip of the auth server")
		inp := ""
		fmt.Scanf("%s", &inp)
		newServ.Ip = net.ParseIP(inp)
		authServBytes, err := bson.Marshal(newServ)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(".dfs.conf")
		if err != nil {
			log.Fatal(err)
		}
		file.Write(authServBytes)
		return newServ
	} else {
		authServ := &authServer{}
		authServBytes, err := ioutil.ReadFile(".dfs.conf")
		if err != nil {
			log.Fatal(err)
		}
		err = bson.Unmarshal(authServBytes, authServ)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(authServ.PubKey.E)
		return authServ
	}
	
}

func main(){

	funcs := make(map[string]func(string))
	authServ := getConfig()
	fmt.Println(authServ.Ip)
	funcs["ls"] = list
	funcs["put"] = put_file
	funcs["transaction start"] = transaction_start
	funcs["transaction end"] = transaction_end
	inp := ""

	fmt.Println("Welcome to DFS")
	auth(authServ)
	help()
	for {
		fmt.Print(">")
		fmt.Scanf("%s", &inp)
		fmt.Println(inp)
		command := funcs[inp]
		if command == nil {
			help()
		} else {
			command(inp)
		}
		
	}
}
