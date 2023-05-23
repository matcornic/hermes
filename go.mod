module github.com/Konstantsiy/hermes/v2

require (
	github.com/Masterminds/sprig v2.16.0+incompatible
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/imdario/mergo v0.3.6
	github.com/jaytaylor/html2text v0.0.0-20180606194806-57d518f124b0
	github.com/matcornic/hermes/v2 v2.1.0
	github.com/russross/blackfriday/v2 v2.0.1
	github.com/stretchr/testify v1.2.2
	golang.org/x/crypto v0.0.0-20181029175232-7e6ffbd03851
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1

go 1.13
