package hermes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testedThemes = []Theme{
	// Insert your new theme here
	new(Default),
}

/////////////////////////////////////////////////////
// Every theme should display the same information //
// Find below the tests to check that              //
/////////////////////////////////////////////////////

// Implement this interface when creating a new example checking a common feature of all themes
type Example interface {
	// Create the hermes example with data
	// Represents the "Given" step in Given/When/Then Workflow
	getExample() (h Hermes, email Email)
	// Checks the content of the generated HTML email by asserting content presence or not
	assertHTMLContent(t *testing.T, s string)
	// Checks the content of the generated Plaintext email by asserting content presence or not
	assertPlainTextContent(t *testing.T, s string)
}

// Scenario
type SimpleExample struct {
	theme Theme
}

func (ed *SimpleExample) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name:      "HermesName",
			Link:      "http://hermes-link.com",
			Copyright: "Copyright © Hermes-Test",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
		},
		TextDirection: TDLeftToRight,
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
			Table: Table{
				Data: [][]Entry{
					{
						{Key: "Item", Value: "Golang"},
						{Key: "Description", Value: "Open source programming language that makes it easy to build simple, reliable, and efficient software"},
						{Key: "Price", Value: "$10.99"},
					},
					{
						{Key: "Item", Value: "Hermes"},
						{Key: "Description", Value: "Programmatically create beautiful e-mails using Golang."},
						{Key: "Price", Value: "$1.99"},
					},
				},
				Columns: Columns{
					CustomWidth: map[string]string{
						"Item":  "20%",
						"Price": "15%",
					},
					CustomAlignement: map[string]string{
						"Price": "right",
					},
				},
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
	return h, email
}

func (ed *SimpleExample) assertHTMLContent(t *testing.T, r string) {

	// Assert on product
	assert.Contains(t, r, "HermesName", "Product: Should find the name of the product in email")
	assert.Contains(t, r, "http://hermes-link.com", "Product: Should find the link of the product in email")
	assert.Contains(t, r, "Copyright © Hermes-Test", "Product: Should find the Copyright of the product in email")
	assert.Contains(t, r, "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png", "Product: Should find the logo of the product in email")

	// Assert on email body
	assert.Contains(t, r, "Hi Jon Snow", "Name: Should find the name of the person")
	assert.Contains(t, r, "Welcome to Hermes", "Intro: Should have intro")
	assert.Contains(t, r, "Birthday", "Dictionary: Should have dictionary")
	assert.Contains(t, r, "Open source programming language", "Table: Should have table with first row and first column")
	assert.Contains(t, r, "Programmatically create beautiful e-mails using Golang", "Table: Should have table with second row and first column")
	assert.Contains(t, r, "$10.99", "Table: Should have table with first row and second column")
	assert.Contains(t, r, "$1.99", "Table: Should have table with second row and second column")
	assert.Contains(t, r, "started with Hermes", "Action: Should have instruction")
	assert.Contains(t, r, "Confirm your account", "Action: Should have button of action")
	assert.Contains(t, r, "#22BC66", "Action: Button should have given color")
	assert.Contains(t, r, "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010", "Action: Button should have link")
	assert.Contains(t, r, "Need help, or have questions", "Outro: Should have outro")
}

func (ed *SimpleExample) assertPlainTextContent(t *testing.T, r string) {

	// Assert on product
	assert.Contains(t, r, "HermesName", "Product: Should find the name of the product in email")
	assert.Contains(t, r, "http://hermes-link.com", "Product: Should find the link of the product in email")
	assert.Contains(t, r, "Copyright © Hermes-Test", "Product: Should find the Copyright of the product in email")
	assert.NotContains(t, r, "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png", "Product: Should not find any logo in plain text")

	// Assert on email body
	assert.Contains(t, r, "Hi Jon Snow", "Name: Should find the name of the person")
	assert.Contains(t, r, "Welcome to Hermes", "Intro: Should have intro")
	assert.Contains(t, r, "Birthday", "Dictionary: Should have dictionary")
	assert.NotContains(t, r, "Open source programming language", "Table: Not possible to have table in plain text")
	assert.NotContains(t, r, "Programmatically create beautiful e-mails using Golang", "Table: Not possible to have table in plain text")
	assert.NotContains(t, r, "$10.99", "Table: Not possible to have table in plain text")
	assert.NotContains(t, r, "$1.99", "Table: Not possible to have table in plain text")
	assert.Contains(t, r, "started with Hermes", "Action: Should have instruction")
	assert.NotContains(t, r, "Confirm your account", "Action: Should not have button of action in plain text")
	assert.NotContains(t, r, "#22BC66", "Action: Button should not have color in plain text")
	assert.Contains(t, r, "https://hermes-example.com/confirm?token=d9729feb74992cc3482b350163a1a010", "Action: Even if button is not possible in plain text, it should have the link")
	assert.Contains(t, r, "Need help, or have questions", "Outro: Should have outro")
}

type WithTitleInsteadOfNameExample struct {
	theme Theme
}

func (ed *WithTitleInsteadOfNameExample) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
	}

	email := Email{
		Body{
			Name:  "Jon Snow",
			Title: "A new e-mail",
		},
	}
	return h, email
}

func (ed *WithTitleInsteadOfNameExample) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Name: should not find greetings from Jon Snow because title should be used")
	assert.Contains(t, r, "A new e-mail", "Title should be used instead of name")
}

func (ed *WithTitleInsteadOfNameExample) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Name: should not find greetings from Jon Snow because title should be used")
	assert.Contains(t, r, "A new e-mail", "Title shoud be used instead of name")
}

type WithGreetingDifferentThanDefault struct {
	theme Theme
}

func (ed *WithGreetingDifferentThanDefault) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
	}

	email := Email{
		Body{
			Greeting: "Dear",
			Name:     "Jon Snow",
		},
	}
	return h, email
}

func (ed *WithGreetingDifferentThanDefault) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Should not find greetings with 'Hi' which is default")
	assert.Contains(t, r, "Dear Jon Snow", "Should have greeting with Dear")
}

func (ed *WithGreetingDifferentThanDefault) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Hi Jon Snow", "Should not find greetings with 'Hi' which is default")
	assert.Contains(t, r, "Dear Jon Snow", "Should have greeting with Dear")
}

type WithSignatureDifferentThanDefault struct {
	theme Theme
}

func (ed *WithSignatureDifferentThanDefault) getExample() (Hermes, Email) {
	h := Hermes{
		Theme: ed.theme,
		Product: Product{
			Name: "Hermes",
			Link: "http://hermes.com",
		},
	}

	email := Email{
		Body{
			Name:      "Jon Snow",
			Signature: "Best regards",
		},
	}
	return h, email
}

func (ed *WithSignatureDifferentThanDefault) assertHTMLContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Yours truly", "Should not find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Best regards", "Should have greeting with Dear")
}

func (ed *WithSignatureDifferentThanDefault) assertPlainTextContent(t *testing.T, r string) {
	assert.NotContains(t, r, "Yours truly", "Should not find signature with 'Yours truly' which is default")
	assert.Contains(t, r, "Best regards", "Should have greeting with Dear")
}

// Test all the themes for the features

func TestThemeSimple(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &SimpleExample{theme})
	}
}

func TestThemeWithTitleInsteadOfName(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithTitleInsteadOfNameExample{theme})
	}
}

func TestThemeWithGreetingDifferentThanDefault(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithGreetingDifferentThanDefault{theme})
	}
}

func TestThemeWithGreetingDiffrentThanDefault(t *testing.T) {
	for _, theme := range testedThemes {
		checkExample(t, &WithSignatureDifferentThanDefault{theme})
	}
}

func checkExample(t *testing.T, ex Example) {
	// Given an example
	h, email := ex.getExample()

	// When generating HTML template
	r, err := h.GenerateHTML(email)
	t.Log(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, r)

	// Then asserting HTML is OK
	ex.assertHTMLContent(t, r)

	// When generating plain text template
	r, err = h.GeneratePlainText(email)
	t.Log(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, r)

	// Then asserting plain text is OK
	ex.assertPlainTextContent(t, r)
}

////////////////////////////////////////////
// Tests on default values for all themes //
// It does not check email content        //
////////////////////////////////////////////

func TestHermes_TextDirectionAsDefault(t *testing.T) {
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
	assert.Equal(t, h.Theme.Name(), "default")
}

func TestHermes_Default(t *testing.T) {
	h := Hermes{}
	setDefaultHermesValues(&h)
	email := Email{}
	setDefaultEmailValues(&email)

	assert.Equal(t, h.TextDirection, TDLeftToRight)
	assert.Equal(t, h.Theme, new(Default))
	assert.Equal(t, h.Product.Name, "Hermes")
	assert.Equal(t, h.Product.Copyright, "Copyright © 2017 Hermes. All rights reserved.")

	assert.Empty(t, email.Body.Actions)
	assert.Empty(t, email.Body.Dictionary)
	assert.Empty(t, email.Body.Intros)
	assert.Empty(t, email.Body.Outros)
	assert.Empty(t, email.Body.Table.Data)
	assert.Empty(t, email.Body.Table.Columns.CustomWidth)
	assert.Empty(t, email.Body.Table.Columns.CustomAlignement)

	assert.Equal(t, email.Body.Greeting, "Hi")
	assert.Equal(t, email.Body.Signature, "Yours truly")
	assert.Empty(t, email.Body.Title)
}
