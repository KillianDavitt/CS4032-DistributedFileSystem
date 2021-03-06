package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/client"
	//"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func help() {
	fmt.Println("\nHelp:\nls\nput\nget")
}

func transaction_start(_ string) {
	fmt.Println("Starting transaction..")
}

func transaction_end(_ string) {
	fmt.Println("End transaction")
}

func login(authServ *auth.AuthServer) []byte {
	username := ""
	password := ""
	fmt.Println("Enter username:")
	fmt.Scanf("%s", &username)
	fmt.Println("Enter password:")
	fmt.Scanf("%s", &password)
	resp, err := authServ.Client.PostForm("https://"+authServ.Ip.String()+":8080/login", url.Values{"username": {username}, "password": {password}})
	if err != nil {
		log.Fatal(err)
	}
	//3ticketMap := auth.GetTicketFromResp(resp.Body, &authServ.PubKey)
	bytes, _ := ioutil.ReadAll(resp.Body)
	return bytes
}

func getDirIp(client *http.Client, ip net.IP) net.IP {
	resp, err := client.Get("https://" + ip.String() + ":8080/get_dir_ip")
	if err != nil {
		log.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	dirIp := net.ParseIP(string(respBytes))

	return dirIp
}

func main() {
	initCache()
	funcs := make(map[string]func([]string, *http.Client, net.IP, []byte))
	funcs["ls"] = client.List
	funcs["put"] = client.Put
	funcs["get"] = client.Get
	//funcs["transaction start"] = transaction_start
	//funcs["transaction end"] = transaction_end
	inp := ""

	fmt.Println("Welcome to DFS")
	authServ := auth.Init()
	ticketMapBytes := login(authServ)
	dirIp := getDirIp(authServ.Client, authServ.Ip)
	scanner := bufio.NewScanner(os.Stdin)
	help()
	for {
		fmt.Print(">")
		scanner.Scan()
		inp = scanner.Text()
		fmt.Println(inp)
		args := strings.Split(inp, " ")
		command := funcs[args[0]]
		if command == nil {
			help()
		} else {
			command(args[1:], authServ.Client, dirIp, ticketMapBytes)
		}

	}
}
