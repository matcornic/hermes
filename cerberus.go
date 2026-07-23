package hermes

// Cerberus is a theme ported from the mailgen "cerberus" theme
// (https://github.com/eladnava/mailgen). It features a dark canvas with a
// white content card and is based on the responsive Cerberus email patterns.
type Cerberus struct{}

// Name returns the name of the cerberus theme
func (dt *Cerberus) Name() string {
	return "cerberus"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Cerberus) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <style type="text/css" rel="stylesheet" media="all">
    /* Base ------------------------------ */
    html, body {
      margin: 0 auto !important;
      padding: 0 !important;
      height: 100% !important;
      width: 100% !important;
    }
    * {
      -ms-text-size-adjust: 100%;
      -webkit-text-size-adjust: 100%;
      font-family: sans-serif;
    }
    table, td {
      mso-table-lspace: 0pt !important;
      mso-table-rspace: 0pt !important;
    }
    img {
      -ms-interpolation-mode: bicubic;
      border: 0;
      height: auto;
      line-height: 100%;
      outline: none;
      text-decoration: none;
    }
    body {
      background-color: #222222;
      color: #555555;
    }
    a {
      color: #222222;
    }
    /* Layout ------------------------------ */
    .email-wrapper {
      width: 100%;
      background-color: #222222;
    }
    .email-card {
      max-width: 600px;
      margin: 0 auto;
      background-color: #ffffff;
    }
    /* Masthead ----------------------- */
    .branding {
      color: #ffffff;
      font-size: 25px;
      font-family: sans-serif;
      padding: 30px 0;
      text-align: center;
    }
    .branding a {
      color: #ffffff;
      text-decoration: none;
    }
    .email-logo {
      max-height: 50px;
    }
    /* Body ------------------------------ */
    .content-cell {
      padding: 10px 40px 40px 40px;
      text-align: center;
      font-family: sans-serif;
      font-size: 15px;
      line-height: 20px;
      color: #555555;
    }
    .content-cell-sub {
      background: #f8f8f8;
      padding: 20px 40px 40px 40px;
      text-align: center;
      font-family: sans-serif;
      font-size: 15px;
      line-height: 20px;
      color: #555555;
    }
    .body-dictionary {
      width: 100%;
      overflow: hidden;
      margin: 0 auto;
      padding: 0;
    }
    .body-dictionary dt {
      clear: both;
      color: #000;
      font-weight: bold;
      margin-right: 4px;
    }
    .body-dictionary dd {
      margin: 0 0 10px 0;
    }
    .body-action {
      width: 100%;
      margin: 30px auto;
      padding: 0;
      text-align: center;
    }
    .body-sub {
      margin-top: 25px;
      padding-top: 25px;
      border-top: 1px solid #EDEFF2;
      table-layout: fixed;
    }
    .body-sub a {
      word-break: break-all;
    }
    .align-right {
      text-align: right;
    }
    /* Type ------------------------------ */
    h1 {
      margin-top: 0;
      color: #333333;
      font-size: 26px;
      font-weight: bold;
      text-align: center;
    }
    h2, h3 {
      margin-top: 0;
      color: #333333;
      font-weight: bold;
    }
    blockquote {
      margin: 25px 0;
      padding-left: 10px;
      border-left: 10px solid #F0F2F4;
    }
    p {
      margin-top: 0;
      color: #555555;
      font-size: 15px;
      line-height: 1.5em;
      text-align: center;
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
      border-bottom: 1px solid #DEDEDE;
    }
    .data-table th p {
      margin: 0;
      color: #9BA2AB;
      font-size: 12px;
      text-align: left;
    }
    .data-table td {
      text-align: left;
      padding: 10px 5px;
      color: #555555;
      font-size: 15px;
      line-height: 18px;
    }
    /* Invite Code ------------------------------ */
    .invite-code {
      display: inline-block;
      padding: 20px 36px 16px 36px;
      border-radius: 3px;
      font-family: Consolas, monaco, monospace;
      font-size: 28px;
      text-align: center;
      letter-spacing: 8px;
      color: #555;
      background-color: #eee;
    }
    /* Buttons ------------------------------ */
    .button {
      display: inline-block;
      background-color: #222222;
      border-radius: 3px;
      color: #ffffff !important;
      font-size: 13px;
      font-weight: bold;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
    .button:hover {
      background: #555555 !important;
    }
    /* Footer ------------------------------ */
    .email-footer {
      max-width: 600px;
      margin: 0 auto;
    }
    .email-footer p {
      color: #888888;
    }
    .copyright a {
      color: #888888;
      text-decoration: underline;
    }
    /*Media Queries ------------------------------ */
    @media only screen and (max-width: 600px) {
      .email-card,
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
<body dir="{{.Hermes.TextDirection}}" width="100%" bgcolor="#222222" style="Margin: 0;">
  <center style="width: 100%; background: #222222;">
    <table class="email-wrapper" width="100%" cellpadding="0" cellspacing="0">
      <tr>
        <td>
          <!-- Logo -->
          <table class="email-card" align="center" width="600" cellpadding="0" cellspacing="0">
            <tr>
              <td class="branding">
                <a href="{{.Hermes.Product.Link}}" target="_blank">
                  {{ if .Hermes.Product.Logo }}
                    <img src="{{.Hermes.Product.Logo | url }}" class="email-logo" />
                  {{ else }}
                    {{ .Hermes.Product.Name }}
                  {{ end }}
                </a>
              </td>
            </tr>
          </table>

          <!-- Email Body -->
          <table class="email-card" align="center" width="600" bgcolor="#ffffff" cellpadding="0" cellspacing="0">
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
                {{ if (ne .Email.Body.FreeMarkdown "") }}
                  {{ .Email.Body.FreeMarkdown.ToHTML }}
                {{ else }}

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
                                      {{ $align := index .CustomAlignment $entry.Key }}
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
                                        {{ $align := index .CustomAlignment $cell.Key }}
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
                        {{ $length := len $action.Button.Text }}
                        {{ $width := add (mul $length 9) 20 }}
                        {{if (lt $width 200)}}{{$width = 200}}{{else if (gt $width 570)}}{{$width = 570}}{{else}}{{end}}
                          {{safe "<!--[if mso]>" }}
                          {{ if $action.Button.Text }}
                            <div style="margin: 30px auto;v-text-anchor:middle;text-align:center">
                              <v:roundrect xmlns:v="urn:schemas-microsoft-com:vml"
                                xmlns:w="urn:schemas-microsoft-com:office:word"
                                href="{{ $action.Button.Link }}"
                                style="height:45px;v-text-anchor:middle;width:{{$width}}px;background-color:{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}#222222{{ end }};"
                                arcsize="10%"
                                {{ if $action.Button.Color }}strokecolor="{{ $action.Button.Color }}" fillcolor="{{ $action.Button.Color }}"{{ else }}strokecolor="#222222" fillcolor="#222222"{{ end }}
                                >
                                <w:anchorlock/>
                                <center style="color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}#FFFFFF{{ end }};font-size: 13px;text-align: center;font-family:sans-serif;font-weight:bold;">
                                  {{ $action.Button.Text }}
                                </center>
                              </v:roundrect>
                            </div>
                          {{ end }}
                          {{ if $action.InviteCode }}
                            <div style="margin-top:30px;margin-bottom:30px">
                              <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
                                <tr>
                                  <td align="center">
                                    <table align="center" cellpadding="0" cellspacing="0" style="padding:0;text-align:center">
                                      <tr>
                                        <td style="display:inline-block;border-radius:3px;font-family:Consolas, monaco, monospace;font-size:28px;text-align:center;letter-spacing:8px;color:#555;background-color:#eee;padding:20px">
                                          {{ $action.InviteCode }}
                                        </td>
                                      </tr>
                                    </table>
                                  </td>
                                </tr>
                              </table>
                            </div>
                          {{ end }}
                          {{safe "<![endif]-->" }}
                          {{safe "<!--[if !mso]><!-- -->"}}
                          <table class="body-action" align="center" width="100%" cellpadding="0" cellspacing="0">
                            <tr>
                              <td align="center">
                                <div>
                                  {{ if $action.Button.Text }}
                                    <a href="{{ $action.Button.Link }}" class="button" style="{{ with $action.Button.Color }}background-color: {{ . }};{{ end }} {{ with $action.Button.TextColor }}color: {{ . }};{{ end }} width: {{$width}}px; padding: 0 25px;" target="_blank">
                                      {{ $action.Button.Text }}
                                    </a>
                                  {{end}}
                                  {{ if $action.InviteCode }}
                                    <span class="invite-code">{{ $action.InviteCode }}</span>
                                  {{end}}
                                </div>
                              </td>
                            </tr>
                          </table>
                          {{safe "<![endif]-->" }}
                      {{ end }}
                    {{ end }}
                  {{ end }}

                {{ end }}
              </td>
            </tr>
            <tr>
              <td class="content-cell">
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
              </td>
            </tr>

            {{ if (eq .Email.Body.FreeMarkdown "") }}
              {{ with .Email.Body.Actions }}
                <tr>
                  <td class="content-cell-sub">
                    <table class="body-sub" width="100%">
                      <tbody>
                          {{ range $action := . }}
                            {{if $action.Button.Text}}
                            <tr>
                              <td>
                                <p class="sub">{{$.Hermes.Product.TroubleText | replace "{ACTION}" $action.Button.Text}}</p>
                                <p class="sub"><a href="{{ $action.Button.Link }}">{{ $action.Button.Link }}</a></p>
                              </td>
                            </tr>
                            {{ end }}
                          {{ end }}
                      </tbody>
                    </table>
                  </td>
                </tr>
              {{ end }}
            {{ end }}
          </table>

          <!-- Footer -->
          <table class="email-footer" align="center" width="600" cellpadding="0" cellspacing="0">
            <tr>
              <td class="copyright" style="padding: 40px 10px; text-align: center;">
                <p class="sub center">
                  {{.Hermes.Product.Copyright}}
                </p>
              </td>
            </tr>
          </table>
        </td>
      </tr>
    </table>
  </center>
</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Cerberus) PlainTextTemplate() string {
	return plainTextTemplate
}
