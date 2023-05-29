module github.com/Konstantsiy/hermes/v2

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	github.com/imdario/mergo v0.3.6
	github.com/jaytaylor/html2text v0.0.0-20180606194806-57d518f124b0
	github.com/matcornic/hermes/v2 v2.1.0
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1
	github.com/stretchr/testify v1.6.1
	github.com/vanng822/go-premailer v1.20.2 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)

replace gopkg.in/russross/blackfriday.v2 v2.0.1 => github.com/russross/blackfriday/v2 v2.0.1

go 1.13
