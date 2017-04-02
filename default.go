package hermes

// Default is the theme by default
type Default struct{}

// Name returns the name of the default theme
func (dt *Default) Name() string {
	return "default"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Default) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <style type="text/css" rel="stylesheet" media="all">
    /* Base ------------------------------ */
    *:not(br):not(tr):not(html) {
      font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif;
      -webkit-box-sizing: border-box;
      box-sizing: border-box;
    }
    body {
      width: 100% !important;
      height: 100%;
      margin: 0;
      line-height: 1.4;
      background-color: #F2F4F6;
      color: #74787E;
      -webkit-text-size-adjust: none;
    }
    a {
      color: #3869D4;
    }
    /* Layout ------------------------------ */
    .email-wrapper {
      width: 100%;
      margin: 0;
      padding: 0;
      background-color: #F2F4F6;
    }
    .email-content {
      width: 100%;
      margin: 0;
      padding: 0;
    }
    /* Masthead ----------------------- */
    .email-masthead {
      padding: 25px 0;
      text-align: center;
    }
    .email-masthead_logo {
      max-width: 400px;
      border: 0;
    }
    .email-masthead_name {
      font-size: 16px;
      font-weight: bold;
      color: #2F3133;
      text-decoration: none;
      text-shadow: 0 1px 0 white;
    }
    .email-logo {
      max-height: 50px;
    }
    /* Body ------------------------------ */
    .email-body {
      width: 100%;
      margin: 0;
      padding: 0;
      border-top: 1px solid #EDEFF2;
      border-bottom: 1px solid #EDEFF2;
      background-color: #FFF;
    }
    .email-body_inner {
      width: 570px;
      margin: 0 auto;
      padding: 0;
    }
    .email-footer {
      width: 570px;
      margin: 0 auto;
      padding: 0;
      text-align: center;
    }
    .email-footer p {
      color: #AEAEAE;
    }
    .body-action {
      width: 100%;
      margin: 30px auto;
      padding: 0;
      text-align: center;
    }
    .body-dictionary {
      width: 100%;
      overflow: hidden;
      margin: 20px auto 10px;
      padding: 0;
    }
    .body-dictionary dd {
      margin: 0 0 10px 0;
    }
    .body-dictionary dt {
      clear: both;
      color: #000;
      font-weight: bold;
    }
    .body-dictionary dd {
      margin-left: 0;
      margin-bottom: 10px;
    }
    .body-sub {
      margin-top: 25px;
      padding-top: 25px;
      border-top: 1px solid #EDEFF2;
    }
    .content-cell {
      padding: 35px;
    }
    .align-right {
      text-align: right;
    }
    /* Type ------------------------------ */
    h1 {
      margin-top: 0;
      color: #2F3133;
      font-size: 19px;
      font-weight: bold;
    }
    h2 {
      margin-top: 0;
      color: #2F3133;
      font-size: 16px;
      font-weight: bold;
    }
    h3 {
      margin-top: 0;
      color: #2F3133;
      font-size: 14px;
      font-weight: bold;
    }
    p {
      margin-top: 0;
      color: #74787E;
      font-size: 16px;
      line-height: 1.5em;
    }
    p.sub {
      font-size: 12px;
    }
    p.center {
      text-align: center;
    }
    /* Data table ------------------------------ */
    .data-wrapper {
      width: 100%;
      margin: 0;
      padding: 35px 0;
    }
    .data-table {
      width: 100%;
      margin: 0;
    }
    .data-table th {
      text-align: left;
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid #EDEFF2;
    }
    .data-table th p {
      margin: 0;
      color: #9BA2AB;
      font-size: 12px;
    }
    .data-table td {
      padding: 10px 5px;
      color: #74787E;
      font-size: 15px;
      line-height: 18px;
    }
    /* Buttons ------------------------------ */
    .button {
      display: inline-block;
      width: 200px;
      background-color: #3869D4;
      border-radius: 3px;
      color: #ffffff;
      font-size: 15px;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
    /*Media Queries ------------------------------ */
    @media only screen and (max-width: 600px) {
      .email-body_inner,
      .email-footer {
        width: 100% !important;
      }
    }
    @media only screen and (max-width: 500px) {
      .button {
        width: 100% !important;
      }
    }
  </style>
</head>
<body dir="{{.Hermes.TextDirection}}">
  <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center">
        <table class="email-content" width="100%" cellpadding="0" cellspacing="0">
          <!-- Logo -->
          <tr>
            <td class="email-masthead">
              <a class="email-masthead_name" href="{{.Hermes.Product.Link}}" target="_blank">
                {{ if .Hermes.Product.Logo }}
                  <img src="{{.Hermes.Product.Logo}}" class="email-logo" />
                {{ else }}
                  {{ .Hermes.Product.Name }}
                {{ end }}
                </a>
            </td>
          </tr>

          <!-- Email Body -->
          <tr>
            <td class="email-body" width="100%">
              <table class="email-body_inner" align="center" width="570" cellpadding="0" cellspacing="0">
                <!-- Body content -->
                <tr>
                  <td class="content-cell">
                    <h1>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h1>
                    {{ with .Email.Body.Intros }}
                      {{ if gt (len .) 0 }}
                        {{ range $line := . }}
                          <p>{{ $line }}</p>
                        {{ end }}
                      {{ end }}
                    {{ end }}

                    {{ with .Email.Body.Dictionary }} 
                      {{ if gt (len .) 0 }}
                        <dl class="body-dictionary">
                          {{ range $entry := . }}
                            <dt>{{ $entry.Key }}:</dt>
                            <dd>{{ $entry.Value }}</dd>
                          {{ end }}
                        </dl>
                      {{ end }}
                    {{ end }}

                    <!-- Table -->
                    {{ with .Email.Body.Table }}
                      {{ $data := .Data }}
                      {{ $columns := .Columns }}
                      {{ if gt (len $data) 0 }}
                        <table class="data-wrapper" width="100%" cellpadding="0" cellspacing="0">
                          <tr>
                            <td colspan="2">
                              <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
                                <tr>
                                  {{ $col := index $data 0 }}
                                  {{ range $entry := $col }}
                                    <th
                                      {{ with $columns }}
                                        {{ $width := index .CustomWidth $entry.Key }}
                                        {{ with $width }}
                                          width="{{ . }}"
                                        {{ end }}
                                        {{ $align := index .CustomAlignement $entry.Key }}
                                        {{ with $align }}
                                          style="text-align:{{ . }}"
                                        {{ end }}
                                      {{ end }}
                                    >
                                      <p>{{ $entry.Key }}</p>
                                    </th>
                                  {{ end }}
                                </tr>
                                {{ range $row := $data }}
                                  <tr>
                                    {{ range $cell := $row }}
                                      <td
                                        {{ with $columns }}
                                          {{ $align := index .CustomAlignement $cell.Key }}
                                          {{ with $align }}
                                            style="text-align:{{ . }}"
                                          {{ end }}
                                        {{ end }}
                                      >
                                        {{ $cell.Value }}
                                      </td>
                                    {{ end }}
                                  </tr>
                                {{ end }}
                              </table>
                            </td>
                          </tr>
                        </table>
                      {{ end }}
                    {{ end }}

                    <!-- Action -->
                    {{ with .Email.Body.Actions }}
                      {{ if gt (len .) 0 }}
                        {{ range $action := . }}
                          <p>{{ $action.Instructions }}</p>
                          <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
                            <tr>
                              <td align="center">
                                <div>
                                  <a href="{{ $action.Button.Link }}" class="button" style="background-color: {{ $action.Button.Color }}" target="_blank">
                                    {{ $action.Button.Text }}
                                  </a>
                                </div>
                              </td>
                            </tr>
                          </table>
                        {{ end }}
                      {{ end }}
                    {{ end }}

                    {{ with .Email.Body.Outros }} 
                      {{ if gt (len .) 0 }}
                        {{ range $line := . }}
                          <p>{{ $line }}</p>
                        {{ end }}
                      {{ end }}
                    {{ end }}

                    <p>
                      {{.Email.Body.Signature}},
                      <br />
                      {{.Hermes.Product.Name}}
                    </p>

                    {{ with .Email.Body.Actions }} 
                      <table class="body-sub">
                        <tbody>
                            {{ range $action := . }}
			      <tr>
                                <td>
                                  <p class="sub">If you’re having trouble with the button '{{ $action.Button.Text }}', copy and paste the URL below into your web browser.</p>
                                  <p class="sub"><a href="{{ $action.Button.Link }}">{{ $action.Button.Link }}</a></p>
                                </td>
			      </tr>
                            {{ end }}
                        </tbody>
                      </table>
                    {{ end }}
                  </td>
                </tr>
              </table>
            </td>
          </tr>
          <tr>
            <td>
              <table class="email-footer" align="center" width="570" cellpadding="0" cellspacing="0">
                <tr>
                  <td class="content-cell">
                    <p class="sub center">
                      {{.Hermes.Product.Copyright}}
                    </p>
                  </td>
                </tr>
              </table>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>
</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Default) PlainTextTemplate() string {
	return `{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }},
{{ with .Email.Body.Intros }}{{ range $line := . }}{{ $line }}{{ end }}{{ end }}
{{ with .Email.Body.Dictionary }}{{ range $entry := . }}
{{ $entry.Key }}: {{ $entry.Value }}{{ end }}{{ end }}
{{ with .Email.Body.Actions }} {{ range $action := . }}
{{ $action.Instructions }}
{{ $action.Button.Link }}{{ end }}{{ end }}
{{ with .Email.Body.Outros }} {{ range $line := . }}
{{ $line }}{{ end }}{{ end }}

{{.Email.Body.Signature}},
{{.Hermes.Product.Name}} - {{.Hermes.Product.Link}}

{{.Hermes.Product.Copyright}}
`
}
