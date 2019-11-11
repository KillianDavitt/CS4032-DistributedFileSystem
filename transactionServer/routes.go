package main

import (
	"github.com/kataras/iris/v12"
	"gopkg.in/redis.v5"
)

func put(ctx iris.Context) {
	// First we need to get the transaction id and make sure that we're in a transaction
	// Otherwise return error

	// Next tell the dirServer we want to put
	//client.Put()

	// Next do a shadow put on the fileServer returned

	// Then we are finished, nothing else happens until the client ends the transaction
	ctx.HTML("Put")
}

func endTransaction(ctx iris.Context) {
	// Get the list of all participating file servers

	// For each, tell them to enter ready to commit

	// If all return true, tell to commit

	// Wait for confirmation

	// Return success to client

	ctx.HTML("Not yet implemented")
}

func beginTransaction(ctx iris.Context) {
	// Ensure no other transaction running

	// Generate a transaction id
	ctx.HTML("Okay")
}

func genTransactionId() string {
	return "id"
}

func getCurrTIDRedis() string {
	client := getTIDRedis()
	tid, err := client.Get("tid").Result()
	if err != nil {
		return "fail"
	}
	return tid
}

func unlockFileRedis(filename string, holder string) bool {
	client := getTIDRedis()
	lock, err := client.Get(filename).Result()
	if err != nil {
		return false
	}
	if lock != holder {
		return false
	} else {
		err = client.Set(filename, "0", 0).Err()
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func getTIDRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 9})
}
