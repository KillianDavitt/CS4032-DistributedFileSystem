// Utils for auth server registering other servers. Auth server stores the ips and keys of all fileservers, dir servers etc
package main

import (
	"crypto/rsa"
	"encoding/json"
	"gopkg.in/redis.v5"
	"log"
	"net"
)

// Enum for types of servers
const (
	DIR = iota
	FILE
	TRANS
)

type Server struct {
	IP     net.IP
	Type   int
	PubKey *rsa.PublicKey
}

func NewServer(ip net.IP, serverType int, pubKey *rsa.PublicKey) *Server {
	newServer := &Server{}
	newServer.IP = ip
	newServer.Type = serverType
	newServer.PubKey = pubKey
	return newServer

}

func ReadServer(ip net.IP) *Server {
	client := getServerRedis()
	serverString, err := client.Get(ip.String()).Result()
	if err != nil {
		log.Fatal(err)
	}
	serverBytes := []byte(serverString)
	var parsedServer Server
	err = json.Unmarshal(serverBytes, &parsedServer)
	if err != nil {
		log.Fatal(err)
	}
	return &parsedServer
}

func (s *Server) writeServerRedis() {
	client := getServerRedis()
	serverBytes, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	serverString := string(serverBytes)
	err = client.Set(s.IP.String(), serverString, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func getServerRedis() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 4})
}

func getServerIps() []net.IP {
	return []net.IP{net.ParseIP("127.0.0.1")}
}

func getServerObjs() []*Server {
	client := getServerRedis()
	keys, err := client.Keys("*").Result()
	if err != nil {
		log.Fatal(err)
	}
	servers := make([]*Server, len(keys), len(keys))
	for i, key := range keys {
		servers[i] = ReadServer(net.ParseIP(key))
	}
	return servers
}

func getDirServers() []*Server {
	servers := getServerObjs()
	dirServers := make([]*Server, 0, 0)
	for _, obj := range servers {
		if obj.Type == DIR {
			dirServers = append(dirServers, obj)
		}
	}
	return dirServers
}

func getDirIps() []net.IP {
	dirServers := getDirServers()
	return []net.IP{dirServers[0].IP}
}
