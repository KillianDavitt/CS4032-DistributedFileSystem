package ticket

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	"time"
)

type ticket struct {
	Token       []byte `json:"genre"`
	Expiry_date time.Time
	Issuee      rsa.PublicKey
}

func genToken() []byte {
	token_len := 32
	token := make([]byte, token_len)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func NewTicket() ticket {
	new_ticket := ticket{}
	new_ticket.Token = genToken()
	//new_ticket.issuee = user
	ticket_length, _ := time.ParseDuration("20h")
	expiry_date := time.Now().Add(ticket_length)
	new_ticket.Expiry_date = expiry_date
	return new_ticket

}

func MarshalTicket(t ticket) []byte {
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
