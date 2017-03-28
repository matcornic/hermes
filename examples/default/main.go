package main

import (
	"fmt"
	"github.com/matcornic/hermes"
	"io/ioutil"
)

type example interface {
	Email() hermes.Email
	Name() string
}

func main() {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Hermes",
			Link: "https://example-hermes.com/",
			Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
	}

	examples := []example{
		new(welcome),
		new(reset),
		new(receipt),
	}

	for _, e := range examples {
		generateEmails(h, e.Email(), e.Name())
	}

}

func generateEmails(h hermes.Hermes, email hermes.Email, example string) {
	// Generate the HTML template and save it
	res, err := h.GenerateHTML(email)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v.%v.html", h.Theme.Name(), example), []byte(res), 0644)
	if err != nil {
		panic(err)
	}

	// Generate the plaintext template and save it
	res, err = h.GeneratePlainText(email)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%v.%v.txt", h.Theme.Name(), example), []byte(res), 0644)
	if err != nil {
		panic(err)
	}
}
