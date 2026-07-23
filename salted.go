package hermes

// Salted is a theme ported from the mailgen "salted" theme
// (https://github.com/eladnava/mailgen), itself based on the responsive
// "Salted" template by A List Apart / Litmus. It is a clean, centered layout
// with alternating white and light-grey sections.
type Salted struct{}

// Name returns the name of the salted theme
func (dt *Salted) Name() string {
	return "salted"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Salted) HTMLTemplate() string {
	return `
<!DOCTYPE html>
<html lang="en" dir="{{.Hermes.TextDirection}}">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width" />
  <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
  <style type="text/css" rel="stylesheet" media="all">
    /* Client-specific ------------------------------ */
    #outlook a { padding: 0; }
    .ReadMsgBody { width: 100%; }
    .ExternalClass { width: 100%; }
    .ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div { line-height: 100%; }
    body, table, td, a { -webkit-text-size-adjust: 100%; -ms-text-size-adjust: 100%; }
    table, td { mso-table-lspace: 0pt; mso-table-rspace: 0pt; }
    img { -ms-interpolation-mode: bicubic; }
    /* Reset ------------------------------ */
    body { margin: 0; padding: 0; height: 100% !important; width: 100% !important; }
    img { border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; }
    body {
      max-width: 575px;
      margin: 0 auto;
      font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
      color: #666666;
    }
    a { color: #68440a; }
    /* Layout ------------------------------ */
    .email-wrapper { width: 100%; }
    .wrapper { width: 500px; margin: 0 auto; }
    /* Masthead ----------------------- */
    .branding {
      font-size: 23px;
      font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
      color: #666666;
      text-decoration: none;
    }
    .branding a {
      color: #666666;
      text-decoration: none;
    }
    .email-logo { max-height: 50px; }
    /* Sections ----------------------- */
    .section-copy {
      padding: 20px 0 0 0;
      font-size: 16px;
      line-height: 25px;
      color: #666666;
      text-align: center;
    }
    /* Type ------------------------------ */
    h1 {
      margin: 0;
      font-size: 25px;
      font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
      color: #333333;
      text-align: center;
    }
    h2, h3 {
      margin: 0;
      color: #333333;
      text-align: center;
    }
    p {
      margin-top: 0;
      color: #666666;
      font-size: 16px;
      line-height: 25px;
      text-align: center;
    }
    p.sub { font-size: 12px; }
    p.center { text-align: center; }
    /* Body ------------------------------ */
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
      margin: 25px auto 0;
      padding: 0;
      text-align: center;
    }
    .body-sub a {
      word-break: break-all;
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
      font-size: 16px;
      line-height: 25px;
      color: #666666;
    }
    .data-table th {
      text-align: left;
      padding: 0px 5px;
      padding-bottom: 8px;
      border-bottom: 1px solid #DEDEDE;
    }
    .data-table th p {
      margin: 0;
      color: #74787E;
      font-size: 12px;
      text-align: left;
    }
    .data-table td {
      padding: 10px 5px;
      color: #666666;
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
      background-color: #48CFAD;
      border-radius: 3px;
      color: #ffffff !important;
      font-size: 16px;
      font-weight: normal;
      line-height: 45px;
      text-align: center;
      text-decoration: none;
      padding: 0 25px;
      -webkit-text-size-adjust: none;
      mso-hide: all;
    }
    /*Media Queries ------------------------------ */
    @media screen and (max-width: 525px) {
      table[class="wrapper"] { width: 100% !important; }
      table[class="responsive-table"] { width: 100% !important; }
      .button { width: 80% !important; }
    }
  </style>
</head>
<body style="margin: 0 auto; padding: 0;" dir="{{.Hermes.TextDirection}}">
  <!-- Header -->
  <table class="email-wrapper" border="0" cellpadding="0" cellspacing="0" width="100%">
    <tr>
      <td bgcolor="#ffffff">
        <div align="center" style="padding: 0px 15px 0px 15px;">
          <table class="wrapper" border="0" cellpadding="0" cellspacing="0" width="500">
            <tr>
              <td style="padding: 30px 0px 30px 0px; text-align: center;">
                <span class="branding">
                  <a href="{{.Hermes.Product.Link}}" target="_blank">
                    {{ if .Hermes.Product.Logo }}
                      <img src="{{.Hermes.Product.Logo | url }}" class="email-logo" />
                    {{ else }}
                      {{ .Hermes.Product.Name }}
                    {{ end }}
                  </a>
                </span>
              </td>
            </tr>
          </table>
        </div>
      </td>
    </tr>
  </table>

  <!-- Main section (light grey) -->
  <table border="0" cellpadding="0" cellspacing="0" width="100%">
    <tr>
      <td bgcolor="#f8f8f8" align="center" style="padding: 70px 15px 70px 15px;">
        <table class="responsive-table" border="0" cellpadding="0" cellspacing="0" width="500">
          <tr>
            <td>
              <h1>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h1>
              <div class="section-copy">
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
                                style="height:45px;v-text-anchor:middle;width:{{$width}}px;background-color:{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}#48CFAD{{ end }};"
                                arcsize="8%"
                                {{ if $action.Button.Color }}strokecolor="{{ $action.Button.Color }}" fillcolor="{{ $action.Button.Color }}"{{ else }}strokecolor="#48CFAD" fillcolor="#48CFAD"{{ end }}
                                >
                                <w:anchorlock/>
                                <center style="color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}#FFFFFF{{ end }};font-size: 16px;text-align: center;font-family:sans-serif;">
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
                                      {{ $action.Button.Text }} &rarr;
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
              </div>
            </td>
          </tr>
        </table>
      </td>
    </tr>
  </table>

  <!-- Lower section (white) -->
  <table border="0" cellpadding="0" cellspacing="0" width="100%">
    <tr>
      <td bgcolor="#ffffff" align="center" style="padding: 15px 15px 40px 15px;">
        <table class="responsive-table" border="0" cellpadding="0" cellspacing="0" width="500">
          <tr>
            <td class="section-copy" style="padding: 25px 30px 0;">
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
                <td style="padding: 0 30px;">
                  <hr />
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
      </td>
    </tr>
    <!-- Footer -->
    <tr>
      <td bgcolor="#f8f8f8" align="center">
        <table width="500" border="0" cellspacing="0" cellpadding="0" align="center" class="responsive-table">
          <tr>
            <td align="center" valign="middle" style="padding: 20px 0px 20px 0px; font-size: 12px; line-height: 18px; color:#666666;">
              <p class="sub center" style="color:#666666;">
                {{.Hermes.Product.Copyright}}
              </p>
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
func (dt *Salted) PlainTextTemplate() string {
	return plainTextTemplate
}
