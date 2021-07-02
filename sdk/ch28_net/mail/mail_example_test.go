package mail

import (
	"fmt"
	"io"
	"log"
	"net/mail"
	"strings"
	"testing"
)

func TestTypeAddress(t *testing.T) {
	addr := mail.Address{
		Name:    "Alice",
		Address: "alice@example.com",
	}
	fmt.Println(addr.String()) // "Alice" <alice@example.com>
}

func TestParseAddress(t *testing.T) {
	addr, err := mail.ParseAddress(`"Alice" <alice@example.com>`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr.Name, addr.Address) // Alice alice@example.com
}

func TestParseAddressList(t *testing.T) {
	const list = "Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>"
	emails, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range emails {
		fmt.Println(v.Name, v.Address)
	}
	// output:
	// Alice alice@example.com
	// Bob bob@example.com
	// Eve eve@example.com
}

func TestTypeMessage(t *testing.T) {
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body
`

	r := strings.NewReader(msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	body, err := io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	// output:
	// Date: Mon, 23 Jun 2015 11:40:36 -0400
	// From: Gopher <from@example.com>
	// To: Another Gopher <to@example.com>
	// Subject: Gophers at Gophercon
	// Message body
}
