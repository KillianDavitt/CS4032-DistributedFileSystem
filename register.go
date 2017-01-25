package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/redis.v5"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Print("Enter username: ")
	var username string
	fmt.Scanln(&username)
	var password string
	fmt.Print("Enter password:")
	fmt.Scanln(&password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	err = client.Set(username, hashedPassword, 0).Err()
	if err != nil {
		panic(err)
	}
}
