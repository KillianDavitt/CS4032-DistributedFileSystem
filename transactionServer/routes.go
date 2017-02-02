package main

import (
	"github.com/kataras/iris"
//	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/client"
)

func get(ctx *iris.Context) {
	ctx.HTML(iris.StatusOK, "Not yet implemented")
}

func put(ctx *iris.Context) {
	// First we need to get the transaction id and make sure that we're in a transaction
	// Otherwise return error


	// Next tell the dirServer we want to put
	//client.Put()

	// Next do a shadow put on the fileServer returned

	// Then we are finished, nothing else happens until the client ends the transaction
}

func endTransaction(ctx *iris.Context) {
	// Get the list of all participating file servers

	// For each, tell them to enter ready to commit

	// If all return true, tell to commit

	// Wait for confirmation

	// Return success to client
	
	ctx.HTML(iris.StatusOK, "Not yet implemented")
}

func beginTransaction(ctx *iris.Context) {
	// Ensure no other transaction running

	// Generate a transaction id
	ctx.HTML(iris.StatusOK, "Okay")
}
