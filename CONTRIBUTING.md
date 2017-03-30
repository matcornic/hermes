# Contributing

Thanks so much for wanting to help! We really appreciate it.

* Have an idea for a new feature?
* Want to add a new built-in theme?

Excellent! You've come to the right place.

1. If you find a bug or wish to suggest a new feature, please create an issue first
2. Make sure your code & comment conventions are in-line with the project's style (execute gometalinter as in [.travis.yml](.travis.yml) file)
3. Make your commits and PRs as tiny as possible - one feature or bugfix at a time
4. Write detailed commit messages, in-line with the project's commit naming conventions

## Theming Instructions

This file contains instructions on adding themes to Mailgen:

* [Using a Custom Theme](#using-a-custom-theme)
* [Creating a Built-In Theme](#creating-a-built-in-theme)

> We use Golang templates under the hood to inject the e-mail body into themes.

### Using a Custom Theme

If you want to supply your own **custom theme** for Hermes to use (but don't want it included with Mailgen):

1. Create a new struct implementing `Theme` interface ([hermes.go](hermes.go)). An real-life example is in [default.go](default.go)
2. Supply your new theme at hermes creation

```go

type MyCustomTheme struct{}

func (dt *MyCustomTheme) Name() string {
	return "mycustomthem"
}

func (dt *MyCustomTheme) HTMLTemplate() string {
    // Get the template from a file (if you want to be able to change the template live without retstarting your application)
    // Or write the template by returning pure string here (if you want embbeded template and do not bother with external dependencies)
    return "<A go html template with wanted information>" 
}

func (dt *Default) PlainTextTemplate() string {
    // Get the template from a file (if you want to be able to change the template live without retstarting your application)
    // Or write the template by returning pure string here (if you want embbeded template and do not bother with external dependencies)
    return "<A go plaintext template with wanter information>"
}

h := hermes.Hermes{
    Theme: new(MyCustomTheme) // Set your fresh new theme here
    Product: hermes.Product{
        Name: "Hermes",
        Link: "https://example-hermes.com/",
    },
}

// ...
// Continue with the rest as usual, create your email and generate the content.
// ...
```

3. That's it.

### Creating a Built-In Theme

If you want to create a new **built-in** Mailgen theme:

1. Fork the repository to your GitHub account and clone it to your computer
2. Create a new Go file named after your new theme
3. Copy content of [default.go](default.go) file in new file and make any necessary changes
6. Scroll down to the [injection snippets](#injection-snippets) and copy and paste each code snippet into the relevant area of your template markup
7. Test the theme by running the `examples/*.js` scripts (insert your theme name in each script) and observing the template output in `preview.html`
8. Take a screenshot of your theme portraying each example and place it in `screenshots/{theme}/{example}.png`
9. Add the theme name, credit, and screenshots to the `README.md` file's [Supported Themes](README.md#supported-themes) section (copy one of the existing themes' markup and modify it accordingly)
7. Submit a pull request with your changes and we'll let you know if anything's missing!

Thanks again for your contribution!

# Injection Snippets

## Product Branding Injection

The following will inject either the product logo or name into the template.

```html
<a href="<%- product.link %>" target="_blank">
    <% if (locals.product.logo) { %>
        <img src="<%- product.logo %>" class="email-logo" />
    <% } else { %>
        <%- product.name %>
    <% } %>
</a>
```

It's a good idea to add the following CSS declaration to set `max-height: 50px` for the logo:

```css
.email-logo {
    max-height: 50px;
}
```

## Title Injection

The following will inject the e-mail title (Hi John Appleseed,) or a custom title provided by the user:

```html
<%- title %>
```

## Intro Injection

The following will inject the intro text (string or array) into the e-mail:

```html
<% if (locals.intro) { %>
    <% intro.forEach(function (introItem) { -%>
        <p><%- introItem %></p>
    <% }) -%>
<% } %>
```

## Dictionary Injection

The following will inject a `<dl>` of key-value pairs into the e-mail:

```html
<!-- Dictionary -->
<% if (locals.dictionary) { %>
    <dl class="dictionary">
    <% for (item in dictionary) { -%>
        <dt><%- item.charAt(0).toUpperCase() + item.slice(1) %>:</dt>
        <dd><%- dictionary[item] %></dd>
    <% } -%>
    </dl>
<% } %>
```

It's a good idea to add this to the top of the template to improve the styling of the dictionary:

```css
/* Dictionary */
.dictionary {
    width: 100%;
    overflow: hidden;
    margin: 0 auto;
    padding: 0;
}
.dictionary dt {
    clear: both;
    color: #000;
    font-weight: bold;
    margin-right: 4px;
}
.dictionary dd {
    margin: 0 0 10px 0;
}
```

## Table Injection

The following will inject the table into the e-mail:

```html
<!-- Table -->
<% if (locals.table) { %>
<table class="data-table" width="100%" cellpadding="0" cellspacing="0">
    <tr>
        <% for (var column in table.data[0]) {%>
        <th
            <% if(locals.table.columns && locals.table.columns.customWidth && locals.table.columns.customWidth[column]) { %>
                width="<%= table.columns.customWidth[column] %>"
            <% } %>
            <% if(locals.table.columns && locals.table.columns.customAlignment && locals.table.columns.customAlignment[column]) { %>
                style="text-align:<%= table.columns.customAlignment[column] %>"
            <% } %>
        >
            <p><%- column.charAt(0).toUpperCase() + column.slice(1) %></p>
        </th>
        <% } %>
    </tr>
    <% for (var i in table.data) {%>
    <tr>
        <% for (var column in table.data[i]) {%>
        <td
            <% if(locals.table.columns && locals.table.columns.customAlignment && locals.table.columns.customAlignment[column]) { %>
                style="text-align:<%= table.columns.customAlignment[column] %>"
            <% } %>
        >
            <%- table.data[i][column] %>
        </td>
        <% } %>
    </tr>
    <% } %>
</table>
<% } %>
```

It's a good idea to add this to the top of the template to improve the styling of the table:

```css
/* Table */
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
    font-size: 12px;
}
.data-table td {
    text-align: left;
    padding: 10px 5px;
    font-size: 15px;
    line-height: 18px;
}
```

## Action Injection

The following will inject the action link (or button) into the e-mail:

```html
<!-- Action -->
<% if (locals.action) { %>
    <% action.forEach(function (actionItem) { -%>
        <p><%- actionItem.instructions %></p>
        <a href="<%- actionItem.button.link %>" style="color:<%- actionItem.button.color %>" target="_blank">
            <%- actionItem.button.text %>
        </a>
    <% }) -%>
<% } %>
```

It's a good idea to add this to the top of the template to specify a fallback color for the action buttons in case it wasn't provided by the user:

```html
<%
if (locals.action) {
    // Make it possible to override action button color (specify fallback color if no color specified)
    locals.action.forEach(function(actionItem) {
        if (!actionItem.button.color) {
            actionItem.button.color = '#48CFAD';
        }
    });
}
%>
```

## Outro Injection

The following will inject the outro text (string or array) into the e-mail:

```html
<% if (locals.outro) { %>
    <% outro.forEach(function (outroItem) { -%>
        <p><%- outroItem %></p>
    <% }) -%>
<% } %>
```

## Signature Injection

The following will inject the signature phrase (e.g. Yours truly) along with the product name into the e-mail:

```html
<%- signature %>,
<br />
<%- product.name %>
```

## Copyright Injection

The following will inject the copyright notice into the e-mail:

```html
<%- product.copyright %>
```

## Go-To Action Injection

In order to support Gmail's [Go-To Actions](https://developers.google.com/gmail/markup/reference/go-to-action), add the following anywhere within the template:

```html
<!-- Support for Gmail Go-To Actions -->
<% if (locals.goToAction) { %>
    <script type="application/ld+json">
    {
        "@context": "http://schema.org",
        "@type": "EmailMessage",
        "potentialAction": {
            "@type": "ViewAction",
            "url": "<%- goToAction.link %>",
            "target": "<%- goToAction.link %>",
            "name": "<%- goToAction.text %>"
        },
        "description": "<%- goToAction.description %>"
    }
    </script>
<% } %>
```

## Text Direction Injection

In order to support generating RTL e-mails, inject the `textDirection` variable into the `<body>` tag:

```html
<body dir="<%- textDirection %>">
```

