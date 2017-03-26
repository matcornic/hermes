package main

import (
	"bytes"
	"github.com/matcornic/hermes"
	"io/ioutil"
)

func main() {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Dictionary: []hermes.Entry{
				{Key: "Firstname", Value: "Jon"},
				{Key: "Lastname", Value: "Snow"},
				{Key: "Birthday", Value: "01/01/283"},
			},
			Actions: []hermes.Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: hermes.Button{
						Color: "#22BC66",
						Text:  "Confirm your account",
						Link:  "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010",
					},
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}

	// Generate the HTML template
	stream, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(stream)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("default.plaintext.html", buf.Bytes(), 0644)

	stream, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}
	buf = new(bytes.Buffer)
	_, err = buf.ReadFrom(stream)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("default.html", buf.Bytes(), 0644)
}
