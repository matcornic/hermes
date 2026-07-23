package hermes

// Neopolitan is a theme ported from the mailgen "neopolitan" theme
// (https://github.com/eladnava/mailgen). It uses the Neapolitan ice-cream
// palette: a teal content area, an orange action area and a dark footer.
type Neopolitan struct{}

// Name returns the name of the neopolitan theme
func (dt *Neopolitan) Name() string {
	return "neopolitan"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Neopolitan) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <style type="text/css" rel="stylesheet" media="all">
    /* Base ------------------------------ */
    @import url(http://fonts.googleapis.com/css?family=Droid+Sans);
    img {
      max-width: 600px;
      outline: none;
      text-decoration: none;
      -ms-interpolation-mode: bicubic;
    }
    a {
      text-decoration: none;
      border: 0;
      outline: none;
      color: #bbbbbb;
    }
    .branding a {
      color: inherit;
    }
    a img {
      border: none;
    }
    td, h1, h2, h3 {
      font-family: 'Droid Sans', Helvetica, Arial, sans-serif;
      font-weight: 400;
    }
    td {
      text-align: center;
    }
    body {
      line-height: 1.7;
      -webkit-font-smoothing: antialiased;
      -webkit-text-size-adjust: none;
      width: 100%;
      height: 100%;
      color: #37302d;
      background: #ffffff;
      font-size: 16px;
    }
    table {
      border-collapse: collapse !important;
    }
    /* Layout ------------------------------ */
    .email-wrapper {
      width: 100%;
      background: #ffffff;
    }
    .email-content {
      margin: 0 auto;
    }
    .force-full-width {
      width: 100% !important;
    }
    /* Masthead ----------------------- */
    .branding {
      font-size: 30px;
      text-align: center;
      padding: 10px 0;
    }
    .email-logo {
      max-height: 50px;
    }
    /* Sections ----------------------- */
    .section-main {
      border-radius: 3px 3px 0 0;
      padding: 0 15px;
    }
    .headline {
      color: #ffffff;
      font-size: 36px;
    }
    .content-cell {
      padding: 0 15px;
    }
    /* Body ------------------------------ */
    .body-dictionary {
      width: 100%;
      overflow: hidden;
      margin: 0 auto 10px;
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
      margin: 20px auto;
      padding: 0;
      text-align: center;
    }
    .body-sub {
      table-layout: fixed;
    }
    .body-sub a {
      word-break: break-all;
    }
    /* Type ------------------------------ */
    h1 {
      margin-top: 0;
      color: #ffffff;
      font-size: 36px;
      font-weight: bold;
    }
    p {
      margin-top: 0;
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
      margin: 0 auto;
    }
    .data-table {
      width: 100%;
      margin: 0 auto;
      font-size: 14px;
    }
    .data-table th {
      text-align: left;
      padding: 0 5px;
      border-bottom: 1px solid rgba(0, 74, 74, 0.3);
    }
    .data-table th p {
      margin: 0;
      font-size: 12px;
      text-align: left;
    }
    .data-table td {
      text-align: left;
      padding: 5px 5px;
      font-size: 14px;
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
      background-color: #414141;
      border-radius: 4px;
      color: #ffffff !important;
      font-size: 16px;
      font-weight: bold;
      line-height: 50px;
      text-align: center;
      text-decoration: none;
      padding: 0 25px;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
    /*Media Queries ------------------------------ */
    @media only screen and (max-width: 480px) {
      table[class="email-content"] {
        width: 320px !important;
      }
    }
  </style>
  <style type="text/css" media="screen">
    @media screen {
      td, h1, h2, h3 {
        font-family: 'Droid Sans', 'Helvetica Neue', 'Arial', 'sans-serif' !important;
      }
    }
  </style>
</head>
<body dir="{{.Hermes.TextDirection}}" style="padding:0; margin:0; display:block; background:#ffffff; -webkit-text-size-adjust:none;" bgcolor="#ffffff">
  <table class="email-wrapper" align="center" cellpadding="0" cellspacing="0" width="100%" height="100%">
    <tr>
      <td align="center" valign="top" bgcolor="#ffffff" width="100%">
        <center>
          <table class="email-content" cellpadding="0" cellspacing="0" width="600">
            <tr>
              <td align="center" valign="top">
                <!-- Logo -->
                <table cellpadding="0" cellspacing="0" width="100%">
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

                <!-- Main section (teal) -->
                <table class="section-main" cellpadding="0" cellspacing="0" width="100%" bgcolor="#4dbfbf">
                  <tr>
                    <td class="headline">
                      <div style="margin-top: 30px">
                        <h1>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h1>
                      </div>
                    </td>
                  </tr>
                  <tr>
                    <td style="color:#003C3C">
                      <br />
                      <table class="content-cell" cellpadding="0" cellspacing="0" width="80%" align="center">
                        <tr>
                          <td style="color:#003C3C">
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
                                            style="height:50px;v-text-anchor:middle;width:{{$width}}px;background-color:{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}#414141{{ end }};"
                                            arcsize="8%"
                                            {{ if $action.Button.Color }}strokecolor="{{ $action.Button.Color }}" fillcolor="{{ $action.Button.Color }}"{{ else }}strokecolor="#414141" fillcolor="#414141"{{ end }}
                                            >
                                            <w:anchorlock/>
                                            <center style="color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}#FFFFFF{{ end }};font-size: 16px;text-align: center;font-family:sans-serif;font-weight:bold;">
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
                                                <a href="{{ $action.Button.Link }}" class="button" style="{{ with $action.Button.Color }}background-color: {{ . }};{{ end }} {{ with $action.Button.TextColor }}color: {{ . }};{{ end }}" target="_blank">
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
                      </table>
                      <br />
                    </td>
                  </tr>
                </table>

                <!-- Lower section (orange) -->
                <table cellpadding="0" cellspacing="0" width="100%" bgcolor="#f5774e">
                  <tr>
                    <td style="color:#801F00">
                      <center>
                        <table cellpadding="0" cellspacing="0" width="80%" align="center">
                          <tr>
                            <td style="color:#801F00">
                              <br />
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

                              {{ if (eq .Email.Body.FreeMarkdown "") }}
                                {{ with .Email.Body.Actions }}
                                  <table class="body-sub" width="100%">
                                    <tbody>
                                        {{ range $action := . }}
                                          {{if $action.Button.Text}}
                                          <tr>
                                            <td>
                                              <p class="sub">{{$.Hermes.Product.TroubleText | replace "{ACTION}" $action.Button.Text}}</p>
                                              <p class="sub"><a href="{{ $action.Button.Link }}" style="color:#801F00">{{ $action.Button.Link }}</a></p>
                                            </td>
                                          </tr>
                                          {{ end }}
                                        {{ end }}
                                    </tbody>
                                  </table>
                                {{ end }}
                              {{ end }}
                              <br />
                            </td>
                          </tr>
                        </table>
                      </center>
                    </td>
                  </tr>
                  <!-- Footer (dark) -->
                  <tr>
                    <td style="background-color:#414141; color:#bbbbbb; padding: 22px 0; font-size:12px;">
                      <p class="sub center" style="color:#bbbbbb;">
                        {{.Hermes.Product.Copyright}}
                      </p>
                    </td>
                  </tr>
                </table>
              </td>
            </tr>
          </table>
        </center>
      </td>
    </tr>
  </table>
</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Neopolitan) PlainTextTemplate() string {
	return plainTextTemplate
}
