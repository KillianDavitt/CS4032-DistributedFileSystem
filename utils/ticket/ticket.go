package ticket

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"
)

type ticket struct {
	token       []byte
	expiry_date time.Time
	issuee      rsa.PublicKey
}

func genToken() []byte {
	token_len := 10
	token := make([]byte, token_len)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func NewTicket() ticket {
	new_ticket := ticket{}
	new_ticket.token = genToken()
	//new_ticket.issuee = user

	ticket_length, _ := time.ParseDuration("20h")
	expiry_date := time.Now().Add(ticket_length)
	new_ticket.expiry_date = expiry_date
	return new_ticket

}
