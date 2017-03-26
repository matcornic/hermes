package hermes

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/imdario/mergo"
	"html/template"
	"io"
)

type TextDirection string

const TDLeftToRight TextDirection = "ltr"
const TDRightToLeft TextDirection = "rtl"

type Hermes struct {
	Theme         Theme
	TextDirection TextDirection // rtl (right to left) or ltr (left to right)
	Product       Product
}

type Product struct {
	// Appears in header & footer of e-mails
	Name      string
	Link      string // e.g. https://matcornic.github.io
	Logo      string // e.g. https://matcornic.github.io/img/logo.png
	Copyright string // Copyright © 2017 Hermes. All rights reserved.
}

type Email struct {
	Body Body
}

type Body struct {
	Name       string
	Intros     []string
	Dictionary []Entry
	Actions    []Action
	Outros     []string
	Greeting   string
	Signature  string
	Title      string
}

type Entry struct {
	Key   string
	Value string
}

type Action struct {
	Instructions string
	Button       Button
}

type Button struct {
	Color string
	Text  string
	Link  string
}

type Template struct {
	Hermes Hermes
	Email  Email
}

type Theme interface {
	Name() string
	HtmlTemplate() string
	PlainTextTemplate() string
}

func setDefaultEmailValues(e *Email) error {
	defaultEmail := Email{
		Body: Body{
			Signature: "Yours truly",
			Greeting:  "Hi",
		},
	}
	return mergo.Merge(e, defaultEmail)
}

func setDefaultHermesValues(h *Hermes) error {
	defaultTextDirection := TDLeftToRight
	defaultHermes := Hermes{
		Theme:         new(DefaultTheme),
		TextDirection: defaultTextDirection,
		Product: Product{
			Name:      "Hermes",
			Copyright: "Copyright © 2017 Hermes. All rights reserved.",
		},
	}
	err := mergo.Merge(h, defaultHermes)
	if err != nil {
		return err
	}
	if h.TextDirection != TDLeftToRight && h.TextDirection != TDRightToLeft {
		h.TextDirection = defaultTextDirection
	}
	return nil
}

func (h *Hermes) GenerateHTML(email Email) (io.Reader, error) {
	return h.generateTemplate(email, func() string {
		return h.Theme.HtmlTemplate()
	})
}

func (h *Hermes) GeneratePlainText(email Email) (io.Reader, error) {
	return h.generateTemplate(email, func() string {
		return h.Theme.PlainTextTemplate()
	})
}

func (h *Hermes) generateTemplate(email Email, tplt func() string) (io.Reader, error) {
	err := setDefaultHermesValues(h)
	if err != nil {
		return nil, err
	}
	setDefaultEmailValues(&email)

	t, err := template.New("hermes").Funcs(sprig.FuncMap()).Parse(tplt())
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	t.Execute(&b, Template{*h, email})
	return &b, nil
}
