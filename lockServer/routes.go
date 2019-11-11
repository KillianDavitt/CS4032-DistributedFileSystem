package main

import (
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
	"github.com/kataras/iris/v12"
	"gopkg.in/redis.v5"
)

func lockFile(ctx iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML("UNAUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	requester := ctx.FormValue("requester")
	succ := lockFileRedis(filename, requester)
	if succ {
		ctx.HTML("Lock granted")
	} else {
		ctx.HTML("Lock is already taken")
	}
}

func unlockFile(ctx iris.Context) {
	if !auth.IsAllowed(ctx) {
		ctx.HTML("UNAUTHORISED")
		return
	}
	filename := ctx.FormValue("filename")
	requester := ctx.FormValue("requester")
	succ := unlockFileRedis(filename, requester)
	if succ {
		ctx.HTML("Lock removed")
	} else {
		ctx.HTML("You are not authorised to remove this lock")
	}
}

func lockFileRedis(filename string, requester string) bool {
	client := getLockRedis()
	lock, err := client.Get(filename).Result()
	if err != nil {
		return false
	}
	if lock != "0" {
		return false
	} else {
		err = client.Set(filename, requester, 0).Err()
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func unlockFileRedis(filename string, holder string) bool {
	client := getLockRedis()
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

func getLockRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 6})
}
