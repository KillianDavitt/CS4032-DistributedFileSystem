package main

import (
	"github.com/kataras/iris"
	//"fmt"
	//"crypto/tls"
)

func main() {
	/*fmt.Println("Please enter the ip of the auth Server...")
	var ipString string
	fmt.Scanln(&ipString)
	fmt.Println(ipString)

	conn, err := tls.Dial("tcp", ipString + ":8080", &tls.Config{})
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("About to print sig")
	fmt.Printf("%x", conn.ConnectionState().PeerCertificates[0].Signature)
*/
	iris.Get("/get_file", getFile)
	iris.Get("/list_files", listFiles)
	iris.Post("/put_file", putFile)
	iris.Get("/register_token", registerToken)
	iris.ListenTLS(":8080", "./cert.pem", "./key_new.pem")
}
