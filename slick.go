package hermes

// Slick is the theme by Konstantin Savchenko https://github.com/konsav/email-templates/
type Slick struct{}

// Name returns name
func (dt *Slick) Name() string {
	return "slick"
}

// HTMLTemplate returns a Golang template that will generate an HTML email.
func (dt *Slick) HTMLTemplate() string {
	return `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8">
  	<meta name="viewport" content="width=device-width, initial-scale=1.0;">
 	<meta name="format-detection" content="telephone=no"/>

	<style>
/* Reset styles */ 
body { margin: 0; padding: 0; min-width: 100%; width: 100% !important; height: 100% !important;}
body, table, td, div, p, a { -webkit-font-smoothing: antialiased; text-size-adjust: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; line-height: 100%; }
table, td { mso-table-lspace: 0pt; mso-table-rspace: 0pt; border-collapse: collapse !important; border-spacing: 0; }
img { border: 0; line-height: 100%; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; }
#outlook a { padding: 0; }
.ReadMsgBody { width: 100%; } .ExternalClass { width: 100%; }
.ExternalClass, .ExternalClass p, .ExternalClass span, .ExternalClass font, .ExternalClass td, .ExternalClass div { line-height: 100%; }

/* Rounded corners for advanced mail clients only */ 
@media all and (min-width: 560px) {
	.container { border-radius: 8px; -webkit-border-radius: 8px; -moz-border-radius: 8px; -khtml-border-radius: 8px;}
}

/* Set color for auto links (addresses, dates, etc.) */ 
a, a:hover {
	color: #127DB3;
}
.footer a, .footer a:hover {
	color: #999999;
}

 	</style>

	<!-- MESSAGE SUBJECT -->
	<title>Get this responsive email template</title>

</head>

<!-- BODY -->
<!-- Set message background color (twice) and text color (twice) -->
<body dir="{{.Hermes.TextDirection}}" topmargin="0" rightmargin="0" bottommargin="0" leftmargin="0" marginwidth="0" marginheight="0" width="100%" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; width: 100%; height: 100%; -webkit-font-smoothing: antialiased; text-size-adjust: 100%; -ms-text-size-adjust: 100%; -webkit-text-size-adjust: 100%; line-height: 100%;
	background-color: #F0F0F0;
	color: #000000;"
	bgcolor="#F0F0F0"
	text="#000000">

<!-- SECTION / BACKGROUND -->
<!-- Set message background color one again -->
<table width="100%" align="center" border="0" cellpadding="0" cellspacing="0" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; width: 100%;" class="background"><tr><td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0;"
	bgcolor="#F0F0F0">

<!-- WRAPPER -->
<!-- Set wrapper width (twice) -->
<table border="0" cellpadding="0" cellspacing="0" align="center"
	width="560" style="border-collapse: collapse; border-spacing: 0; padding: 0; width: inherit;
	max-width: 560px;" class="wrapper">

	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 20px;
			padding-bottom: 20px;">

			<!-- PREHEADER -->
			<!-- Set text color to background color -->
			<div style="display: none; visibility: hidden; overflow: hidden; opacity: 0; font-size: 1px; line-height: 1px; height: 0; max-height: 0; max-width: 0;
			color: #F0F0F0;" class="preheader">
				Available on&nbsp;GitHub and&nbsp;CodePen. Highly compatible. Designer friendly. More than 50%&nbsp;of&nbsp;total email opens occurred on&nbsp;a&nbsp;mobile device&nbsp;— a&nbsp;mobile-friendly design is&nbsp;a&nbsp;must for&nbsp;email campaigns.</div>

			<!-- LOGO -->
			<!-- Image text color should be opposite to background color. Set your url, image src, alt and title. Alt text should fit the image size. Real image size should be x2. URL format: http://domain.com/?utm_source={{Campaign-Source}}&utm_medium=email&utm_content=logo&utm_campaign={{Campaign-Name}} -->
			<a target="_blank" style="text-decoration: none;" href="{{.Hermes.Product.Link}}">
                {{ if .Hermes.Product.Logo }}
                  <img border="0" vspace="0" hspace="0"
                  src="{{.Hermes.Product.Logo | url }}"
                  alt="Logo" title="Logo" style="
                  color: #000000;
                  font-size: 10px; margin: 0; padding: 0; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; border: none; display: block;" />
                {{ else }}
                  {{ .Hermes.Product.Name }}
                {{ end }}
            </a>

		</td>
	</tr>

<!-- End of WRAPPER -->
</table>

<!-- WRAPPER / CONTEINER -->
<!-- Set conteiner background color -->
<table border="0" cellpadding="0" cellspacing="0" align="center"
	bgcolor="#FFFFFF"
	width="560" style="border-collapse: collapse; border-spacing: 0; padding: 0; width: inherit;
	max-width: 560px;" class="container">

	<!-- HEADER -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 24px; font-weight: bold; line-height: 130%;
			padding-top: 25px;
			color: #000000;
			font-family: sans-serif;" class="header">
                {{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}
		</td>
	</tr>
	
	<!-- SUBHEADER -->
    {{ with .Email.Body.Intros }}
        {{ if gt (len .) 0 }}
          {{ range $line := . }}
          <tr>
              <td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-bottom: 3px; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 18px; font-weight: 300; line-height: 150%;
                  padding-top: 5px;
                  color: #000000;
                  font-family: sans-serif;" class="subheader">
                  {{ $line }}
              </td>
          </tr>
          {{ end }}
        {{ end }}
    {{ end }}
    {{ if (ne .Email.Body.FreeMarkdown "") }}
      {{ .Email.Body.FreeMarkdown.ToHTML }}
    {{ else }}
    {{ if (ne .Email.Body.Hero.URL "") }}
	<!-- HERO IMAGE -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 10;
			padding-top: 20px;" class="hero"><a target="_blank" style="text-decoration: none;"
			href="{{.Hermes.Product.Link}}"><img border="0" vspace="0" hspace="0"
            src="{{.Email.Body.Hero.URL}}"
			alt="Please enable images to view this content" title="{{.Email.Body.Hero.Title}}"
			width="{{.Email.Body.Hero.Width}}" height="{{.Email.Body.Hero.Height}}" style="
			width: 100%;
			max-width: 560px;
			color: #000000; font-size: 13px; margin: 0; padding: 0; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; border: none; display: block;"/></a></td>
	</tr>
    {{ end }}

    {{ with .Email.Body.Dictionary }} 
      {{ if gt (len .) 0 }}
	<tr>	
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;" class="paragraph">
        <dl>
          {{ range $entry := . }}
            <dt>{{ $entry.Key }}:</dt>
            <dd>{{ $entry.Value }}</dd>
          {{ end }}
        </dl>
        </td>
	</tr>
      {{ end }}
    {{ end }}

    {{ with .Email.Body.Exploration }} 
      {{ if gt (len .) 0 }}
	<!-- LINE -->
	<!-- Set line color -->
	<tr>	
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;" class="line"><hr
			color="#E0E0E0" align="center" width="100%" size="1" noshade style="margin: 0; padding: 0;" />
		</td>
	</tr>

          {{ range $entry := . }}
	<!-- LIST -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%;" class="list-item"><table align="center" border="0" cellspacing="0" cellpadding="0" style="width: inherit; margin: 0; padding: 0; border-collapse: collapse; border-spacing: 0;">
			
			<!-- LIST ITEM -->
			<tr>

				<!-- LIST ITEM IMAGE -->
				<!-- Image text color should be opposite to background color. Set your url, image src, alt and title. Alt text should fit the image size. Real image size should be x2 -->
				<td align="left" valign="top" style="border-collapse: collapse; border-spacing: 0;
					padding-top: 30px;
					padding-right: 20px;"><img
				border="0" vspace="0" hspace="0" style="border-radius: 50%; padding: 0; margin: 0; outline: none; text-decoration: none; -ms-interpolation-mode: bicubic; border: none; display: block;
					color: #000000;"
					src="{{ $entry.Image.URL }}"
					alt="H" title="{{ $entry.Image.Title }}"
					width="{{ $entry.Image.Width }}" height="{{ $entry.Image.Height }}"></td>

				<!-- LIST ITEM TEXT -->
				<!-- Set text color and font family ("sans-serif" or "Georgia, serif"). Duplicate all text styles in links, including line-height -->
				<td align="left" valign="top" style="font-size: 17px; font-weight: 400; line-height: 160%; border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0;
					padding-top: 25px;
					color: #000000;
					font-family: sans-serif;" class="paragraph">
						<b style="color: #333333;">{{ $entry.Title }}</b><br/>
						{{ $entry.Description }}
				</td>

			</tr>

		</table></td>
	</tr>
          {{ end }}

	<!-- LINE -->
	<!-- Set line color -->
	<tr>	
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;" class="line"><hr
			color="#E0E0E0" align="center" width="100%" size="1" noshade style="margin: 0; padding: 0;" />
		</td>
	</tr>
      {{ end }}
    {{ end }}

                      <!-- Table -->
                      {{ with .Email.Body.Table }}
                        {{ $data := .Data }}
                        {{ $columns := .Columns }}
                        {{ if gt (len $data) 0 }}
	<tr>	
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;" class="paragraph">
                          <table width="100%" cellpadding="0" cellspacing="0">
                            <tr>
                              <td colspan="2">
                                <table width="100%" cellpadding="0" cellspacing="0">
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
		</td>
	</tr>
    {{ end }}
  {{ end }}

  <!-- Action -->
  {{ with .Email.Body.Actions }}
    {{ if gt (len .) 0 }}
      {{ range $action := . }}
	<!-- PARAGRAPH -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 17px; font-weight: 400; line-height: 160%;
			padding-top: 25px; 
			color: #000000;
			font-family: sans-serif;" class="paragraph">
                {{ $action.Instructions }}
		</td>
	</tr>

        {{ if $action.Button.Text }}
	<!-- BUTTON -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;
			padding-bottom: 55px;" class="button">
				<table border="0" cellpadding="0" cellspacing="0" align="center" style="max-width: 240px; min-width: 120px; border-collapse: collapse; border-spacing: 0; padding: 0;"><tr><td align="center" valign="middle" style="padding: 12px 24px; margin: 0; text-decoration: underline; border-collapse: collapse; border-spacing: 0; border-radius: 4px; -webkit-border-radius: 4px; -moz-border-radius: 4px; -khtml-border-radius: 4px;"
					bgcolor="{{ if $action.Button.Color }}{{ $action.Button.Color }}{{ else }}#E9703E{{ end }}"><a target="_blank" style="text-decoration: underline;
					color: {{ if $action.Button.TextColor }}{{ $action.Button.TextColor }}{{else}}#FFFFFF{{ end }}; font-family: sans-serif; font-size: 17px; font-weight: 400; line-height: 120%;"
					href="{{ $action.Button.Link }}">
                        {{ $action.Button.Text }}
					</a>
			</td></tr></table>
		</td>
	</tr>
        {{ end }}
        {{ if $action.InviteCode }}
	<!-- CODE -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%;
			padding-top: 25px;
			padding-bottom: 55px;" class="button">
				<table border="0" cellpadding="0" cellspacing="0" align="center" style="max-width: 240px; min-width: 120px; border-collapse: collapse; border-spacing: 0; padding: 0;"><tr><td align="center" valign="middle" style="padding: 12px 24px; margin: 0; text-decoration: underline; border-collapse: collapse; border-spacing: 0; border-radius: 4px; -webkit-border-radius: 4px; -moz-border-radius: 4px; -khtml-border-radius: 4px; color: #555; font-family: sans-serif; font-size: 28px; font-weight: 400; line-height: 120%;" bgcolor="#eee">
                        {{ $action.InviteCode }}
			</td></tr></table>
		</td>
	</tr>
        {{ end }}
      {{ end }}
    {{ end }}
  {{ end }}

    {{ with .Email.Body.Outros }} 
      {{ if gt (len .) 0 }}
        {{ range $line := . }}
	<!-- PARAGRAPH -->
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 17px; font-weight: 400; line-height: 160%;
			padding-top: 25px; 
			color: #000000;
			font-family: sans-serif;" class="paragraph">
             {{ $line }}
		</td>
	</tr>
        {{ end }}
      {{ end }}
    {{ end }}

<!-- End of WRAPPER -->
</table>

<!-- WRAPPER -->
<!-- Set wrapper width (twice) -->
<table border="0" cellpadding="0" cellspacing="0" align="center"
	width="560" style="border-collapse: collapse; border-spacing: 0; padding: 0; width: inherit;
	max-width: 560px;" class="wrapper">

	<!-- FOOTER -->
                    {{ if (eq .Email.Body.FreeMarkdown "") }}
                      {{ with .Email.Body.Actions }} 
                              {{ range $action := . }}
                                {{if $action.Button.Text}}
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 13px; font-weight: 400; line-height: 150%;
			padding-top: 20px;
			padding-bottom: 0px;
			color: #555;
			font-family: sans-serif;" class="footer">
                {{$.Hermes.Product.TroubleText | replace "{ACTION}" $action.Button.Text}}
				<a href="{{ $action.Button.Link }}" target="_blank" style="text-decoration: underline; color: blue; font-family: sans-serif; font-size: 13px; font-weight: 400; line-height: 150%;">{{ $action.Button.Link }}</a>
		</td>
	</tr>
                                {{ end }}
                              {{ end }}
                      {{ end }}
                    {{ end }}
    {{ if (ne .Email.Body.Unsubscribe.URL "") }}
	<tr>
		<td align="center" valign="top" style="border-collapse: collapse; border-spacing: 0; margin: 0; padding: 0; padding-left: 6.25%; padding-right: 6.25%; width: 87.5%; font-size: 13px; font-weight: 400; line-height: 150%;
			padding-top: 20px;
			padding-bottom: 20px;
			color: #999999;
			font-family: sans-serif;" class="footer">

            {{ .Email.Body.Unsubscribe.Text }} <a href="{{ .Email.Body.Unsubscribe.URL }}" target="_blank" style="text-decoration: underline; color: #999999; font-family: sans-serif; font-size: 13px; font-weight: 400; line-height: 150%;">{{ .Email.Body.Unsubscribe.Anchor }}</a>.
		</td>
	</tr>
    {{ end }}

<!-- End of WRAPPER -->
</table>

<!-- End of SECTION / BACKGROUND -->
</td></tr></table>

</body>
</html>
`
}

// PlainTextTemplate returns a Golang template that will generate an plain text email.
func (dt *Slick) PlainTextTemplate() string {
	return `<h2>{{if .Email.Body.Title }}{{ .Email.Body.Title }}{{ else }}{{ .Email.Body.Greeting }} {{ .Email.Body.Name }},{{ end }}</h2>
{{ with .Email.Body.Intros }}
  {{ range $line := . }}
    <p>{{ $line }}</p>
  {{ end }}
{{ end }}
{{ if (ne .Email.Body.FreeMarkdown "") }}
  {{ .Email.Body.FreeMarkdown.ToHTML }}
{{ else }}
  {{ with .Email.Body.Dictionary }}
    <ul>
    {{ range $entry := . }}
      <li>{{ $entry.Key }}: {{ $entry.Value }}</li>
    {{ end }}
    </ul>
  {{ end }}
  {{ with .Email.Body.Table }}
    {{ $data := .Data }}
    {{ $columns := .Columns }}
    {{ if gt (len $data) 0 }}
      <table class="data-table" width="100%" cellpadding="0" cellspacing="0">
        <tr>
          {{ $col := index $data 0 }}
          {{ range $entry := $col }}
            <th>{{ $entry.Key }} </th>
          {{ end }}
        </tr>
        {{ range $row := $data }}
          <tr>
            {{ range $cell := $row }}
              <td>
                {{ $cell.Value }}
              </td>
            {{ end }}
          </tr>
        {{ end }}
      </table>
    {{ end }}
  {{ end }}
  {{ with .Email.Body.Exploration }} 
    {{ if gt (len .) 0 }}
------------------------------------------
        {{ range $entry := . }}
## {{ $entry.Title }}
{{ $entry.Description }}

        {{ end }}
==========================================
    {{ end }}
  {{ end }}
  {{ with .Email.Body.Actions }} 
    {{ range $action := . }}
      <p>
        {{ $action.Instructions }} 
        {{ if $action.InviteCode }}
          {{ $action.InviteCode }}
        {{ end }}
        {{ if $action.Button.Link }}
          {{ $action.Button.Link }}
        {{ end }}
      </p> 
    {{ end }}
  {{ end }}
{{ end }}
{{ with .Email.Body.Outros }} 
  {{ range $line := . }}
    <p>{{ $line }}<p>
  {{ end }}
{{ end }}
<p>{{.Email.Body.Signature}},<br>{{.Hermes.Product.Name}} - {{.Hermes.Product.Link}}</p>

<p>{{.Hermes.Product.Copyright}}</p>
`
}
