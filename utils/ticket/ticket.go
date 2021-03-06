package ticket

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/KillianDavitt/CS4032-DistributedFileSystem/utils/rsaUtil"
	"log"
	"time"
)

type Ticket struct {
	Token       string `json:"genre"`
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

func NewTicket() Ticket {
	new_ticket := Ticket{}
	new_ticket.Token = string(genToken())
	//new_ticket.issuee = user
	ticket_length, _ := time.ParseDuration("20h")
	expiry_date := time.Now().Add(ticket_length)
	new_ticket.Expiry_date = expiry_date
	return new_ticket

}

func (t Ticket) MarshalTicket() []byte {
	data, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (t Ticket) CreateTicketMap(privKey *rsa.PrivateKey) string {
	ticketMap := make(map[string][]byte)
	ticketBytes := t.MarshalTicket()
	ticketMap["ticket"] = ticketBytes
	signedTicket := rsaUtil.Sign(ticketBytes, privKey)
	ticketMap["signed_ticket"] = signedTicket
	jsonBytes, err := json.Marshal(ticketMap)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

func GetTicketMap(ticketMapString string, pubKey *rsa.PublicKey) Ticket {
	ticketMap := make(map[string][]byte)
	fmt.Println(ticketMapString)
	err := json.Unmarshal([]byte(ticketMapString), &ticketMap)
	if err != nil {
		log.Fatal(err)
	}

	providedSig := ticketMap["signed_ticket"]
	ticketData := ticketMap["ticket"]
	validSignature := rsaUtil.Verify(pubKey, ticketData, providedSig)
	if !validSignature {
		log.Fatal("Failure verifying ticket signature, possible MITM!")
	}

	var newTicket Ticket
	err = json.Unmarshal(ticketData, &newTicket)
	if err != nil {
		log.Fatal(err)
	}
	return newTicket
}
