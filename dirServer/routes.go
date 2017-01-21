package main
import (
	"github.com/kataras/iris"
	"log"
	"gopkg.in/redis.v5"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/auth"
)
func getFile(ctx *iris.Context){
	//filename := ctx.URLParam("filename")
	//f := readFiles()
	//ip := f.getFile(filename).String()
	ctx.HTML(iris.StatusOK, "hi")
}

func putFile(ctx *iris.Context){
	ctx.HTML(iris.StatusOK, "ok")
}


func listFiles(ctx *iris.Context){
	//	f := readFiles()
	ctx.HTML(iris.StatusOK, "Hi")//f.getFile("test.txt").String())
}

func registerToken(ctx *iris.Context){
	pubKey := auth.RetrieveKey("authserver")
	token := ctx.Param("token")
	ticket := ticket.GetTicketMap(token, pubKey)
	token_client := getTokenRedis()
	expiryString, err := ticket.Expiry_date.MarshalText()
	if err != nil {
		log.Fatal(err)
	}
	err = token_client.Set(string(ticket.Token), string(expiryString), 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func isAllowed(token []byte)(bool){
	client := getTokenRedis()
	_, err := client.Get(string(token)).Result()
	if err != nil{
		return false
	}
	return true
}

func getTokenRedis() (*redis.Client){
	return redis.NewClient(&redis.Options{ Addr: "localhost:6379", Password: "", DB: 1})
}
