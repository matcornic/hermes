package hermes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHermes(t *testing.T) {

	h := Hermes{
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Dictionary: []Entry{
				{"Firstname", "Jon"},
				{"Lastname", "Snow"},
				{"Birthday", "01/01/283"},
			},
			Actions: []Action{
				{
					Instructions: "To get started with Hermes, please click here:",
					Button: Button{
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

	r, err := h.GenerateHTML(email)
	t.Log(r)
	assert.Nil(t, err)
}
