package auth

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/ticket"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type AuthServer struct {
	Ip     net.IP
	Cert   x509.Certificate
	PubKey rsa.PublicKey
	Client *http.Client
}

func writeConfig(authServ *AuthServer) {
	fmt.Println(authServ.PubKey.E)
	authServBytes, err := bson.Marshal(authServ)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(".dfs.conf")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(authServBytes)
}

func getConfig() *AuthServer {
	if _, err := os.Stat(".dfs.conf"); os.IsNotExist(err) {
		newServ := &AuthServer{}
		fmt.Println("Enter the ip of the auth server")
		inp := ""
		fmt.Scanf("%s", &inp)
		newServ.Ip = net.ParseIP(inp)
		authServBytes, err := bson.Marshal(newServ)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(".dfs.conf")
		if err != nil {
			log.Fatal(err)
		}
		file.Write(authServBytes)
		return newServ
	} else {
		authServ := &AuthServer{}
		authServBytes, err := ioutil.ReadFile(".dfs.conf")
		if err != nil {
			log.Fatal(err)
		}
		err = bson.Unmarshal(authServBytes, authServ)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(authServ.PubKey.E)
		return authServ
	}

}

func Init() *AuthServer {
	authServ := getConfig()

	// InsecureSkipVerify must be set since we need to contact the auth server once to find it's fingerprint
	conn, err := tls.Dial("tcp", authServ.Ip.String()+":8080", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	servedCert := *conn.ConnectionState().PeerCertificates[0]
	servedPubKey := rsa.PublicKey(*servedCert.PublicKey.(*rsa.PublicKey))

	fmt.Println(authServ.PubKey.E)
	fmt.Println(servedPubKey.E)

	if authServ.PubKey == servedPubKey {
		fmt.Println("Keys match")
	} else {
		fmt.Println("You have not saved this auth servers public key...")
		fmt.Println(servedPubKey.E)
		fmt.Println("Would you like to accept this key?")
		input := ""
		fmt.Scanf("%s", &input)
		if input == "y" {
			authServ.PubKey = servedPubKey
			authServ.Cert = servedCert
			writeConfig(authServ)
			StoreRedis(&servedPubKey, "authserver")
			fmt.Println("This auth server public key has been accepted")
		}
	}
	client := GetClientFromCert(&servedCert)
	authServ.Client = client
	return authServ
}

func LoadCertFromDisk(filename string) *x509.Certificate {
	certBytes, _ := ioutil.ReadFile(filename)
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	return cert
}

func GetClientFromCert(cert *x509.Certificate) *http.Client {
	CA_Pool := x509.NewCertPool()
	CA_Pool.AddCert(cert)

	// This bit set to insecure until I can fix it later
	tlsConf := &tls.Config{RootCAs: CA_Pool, InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: transport}
	return client
}

func GetTicketFromResp(body io.Reader, pubKey *rsa.PublicKey) ticket.Ticket {
	bytes, _ := ioutil.ReadAll(body)
	parsedTicketMap := ticket.GetTicketMap(string(bytes), pubKey)
	return parsedTicketMap
}
