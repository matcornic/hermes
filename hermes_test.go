package hermes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHermes_ok(t *testing.T) {

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
	assert.NotEmpty(t, r)

	r, err = h.GeneratePlainText(email)
	t.Log(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, r)

	assert.Equal(t, h.Theme.Name(), "default")
}

func TestHermes_defaultTextDirection(t *testing.T) {
	h := Hermes{
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
		TextDirection: "not-existing", // Wrong text-direction
	}

	email := Email{
		Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
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

	_, err := h.GenerateHTML(email)
	assert.Nil(t, err)
	assert.Equal(t, h.TextDirection, TDLeftToRight)
}
