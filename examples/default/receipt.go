package main

import (
	"github.com/matcornic/hermes"
)

type Receipt struct {
}

func (r *Receipt) Name() string {
	return "receipt"
}

func (r *Receipt) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Your order has been processed successfully.",
			},
			Table: hermes.Table{
				Data: [][]hermes.Entry{
					{
						{Key: "Item", Value: "Node.js"},
						{Key: "Description", Value: "Event-driven I/O server-side JavaScript environment based on V8"},
						{Key: "Price", Value: "$10.99"},
					},
					{
						{Key: "Item", Value: "Mailgen"},
						{Key: "Description", Value: "Programmatically create beautiful e-mails using plain old JavaScript."},
						{Key: "Price", Value: "$1.99"},
					},
				},
				Columns: hermes.Columns{
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignement: map[string]string{
						"Price": "right",
					},
				},
			},
			Actions: []hermes.Action{
				{
					Instructions: "You can check the status of your order and more in your dashboard:",
					Button: hermes.Button{
						Text: "Go to Dashboard",
						Link: "https://hermes-example.com/dashboard",
					},
				},
			},
		},
	}
}
