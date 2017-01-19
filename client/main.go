package main

import (
	"net"
	"crypto/rsa"
	"fmt"
	"os"
	"gopkg.in/mgo.v2/bson"
	"log"
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

func auth(_ string){
	fmt.Println("Authing...")
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

func getConfig(){
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
	}
}

func main(){

	funcs := make(map[string]func(string))
	getConfig()
	funcs["ls"] = list
	funcs["put"] = put_file
	funcs["transaction start"] = transaction_start
	funcs["transaction end"] = transaction_end
	inp := ""

	fmt.Println("Welcome to DFS")
	auth("hi")
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
