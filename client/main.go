package main

import (
	"fmt"
)

func help(){
	fmt.Println("Unknown command")
}

func list(_ string){
	fmt.Println("Listing")
}

func main(){

	funcs := make(map[string]func(string))

	funcs["ls"] = list
	
	inp := ""
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
