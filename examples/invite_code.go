package main

import (
	"github.com/matcornic/hermes/v2"
)

type inviteCode struct {
}

func (w *inviteCode) Name() string {
	return "invite_code"
}

func (w *inviteCode) Email() hermes.Email {
	return hermes.Email{
		Body: hermes.Body{
			Name: "Jon Snow",
			Intros: []string{
				"Welcome to Hermes! We're very excited to have you on board.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Please copy your invite code:",
					InviteCode: "123456",
				},
			},
			Outros: []string{
				"Need help, or have questions? Just reply to this email, we'd love to help.",
			},
		},
	}
}
