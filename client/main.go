package main

import (
	"fmt"
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

func main(){

	funcs := make(map[string]func(string))
	
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
