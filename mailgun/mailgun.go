package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	mailgun "github.com/mailgun/mailgun-go"
)

func main() {
	domain := "SOURCE_DOMAIN"
	apiKey := "APIKEY"
	source := "source@domain.com"
	destination := "destination@domain.com"

	resp, err := SendSimpleMessage(domain, apiKey, source, destination)
	if err != nil {
		message := fmt.Sprintf("Error sending email %q", err)
		panic(message)
	}

	fmt.Printf("\n%s", resp)
}

// SendSimpleMessage Invoke mailgun
func SendSimpleMessage(domain, apiKey, source, destination string) (string, error) {
	b, err := ioutil.ReadFile("template.html")
	if err != nil {
		fmt.Print(err)
	}

	mg := mailgun.NewMailgun(domain, apiKey)
	m := mg.NewMessage(
		source,
		"Sending Html Test",
		"Testing some Mailgun Awesomeness!",
		destination,
	)
	m.SetHtml(string(b))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := mg.Send(ctx, m)
	return id, err
}
