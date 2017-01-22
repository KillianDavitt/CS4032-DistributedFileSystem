package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"fmt"
	"log"
	"net/http"
	"net"
	"net/url"
)

func help(){
	fmt.Println("Help:\nls\nput\nget")
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

func login(client *http.Client, ip net.IP){
	username := ""
	password := ""
	fmt.Println("Enter username:")
	fmt.Scanf("%s", &username)
	fmt.Println("Enter password:")
	fmt.Scanf("%s", &password)
	resp, err := client.PostForm("https://" + ip.String() + ":8080/login", url.Values{"username": {username}, "password": {password}}) 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func getDirIp(client *http.Client, ip net.IP) (net.IP){
	_, err := client.Get("https://" + ip.String() + ":8080/get_dir_ip")
	if err != nil {
		log.Fatal(err)
	}
	return net.ParseIP("0.0.0.0")
}

func main(){

	funcs := make(map[string]func(string, *http.Client, net.IP))
	funcs["ls"] = list
	//funcs["put"] = put_file
	//funcs["transaction start"] = transaction_start
	//funcs["transaction end"] = transaction_end
	inp := ""

	fmt.Println("Welcome to DFS")
	conn, ip := auth.Init()
	login(conn, ip)
	dirIp := getDirIp(conn, ip)
	
	help()
	for {
		fmt.Print(">")
		fmt.Scanf("%s", &inp)
		fmt.Println(inp)
		command := funcs[inp]
		if command == nil {
			help()
		} else {
			command(inp, conn, dirIp)
		}
		
	}
}
