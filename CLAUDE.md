# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Hermes is a Go library (package `hermes`, import path `github.com/matcornic/hermes`) that generates clean, responsive HTML transactional emails plus a plain-text fallback. It is a Go port of the Node.js [mailgen](https://github.com/eladnava/mailgen) engine.

## Commands

```bash
go build ./...                 # build
go test ./...                  # run all tests
go test -run TestThemeXXX      # run a single test by name
golangci-lint run              # lint (config in .golangci.yml; CI runs `-v`, only-new-issues)
```

Run the examples (writes HTML + plaintext files into `examples/default` and `examples/flat`):

```bash
cd examples && go run -a *.go
```

Examples can also send real emails for visual testing in email clients by setting `HERMES_SEND_EMAILS=true` plus the `HERMES_SMTP_*` / `HERMES_SENDER_*` / `HERMES_TO` env vars (see README "More Examples").

## Architecture

The public API lives in **`hermes.go`**. A user builds a `Hermes` engine (theme + product/brand info + text direction) and an `Email` (a `Body` of intros, actions, tables, dictionaries, outros, etc.), then calls:

- `GenerateHTML(email)` → inlined-CSS HTML for modern clients.
- `GeneratePlainText(email)` → HTML rendered then converted to text via `html2text`.

Both funnel into `generateTemplate`, the core pipeline:

1. `setDefaultHermesValues` / `setDefaultEmailValues` fill zero-value fields using `mergo.Merge` (defaults never override user-set values).
2. The theme template string is parsed with Go `html/template`, extended with [sprig](https://github.com/Masterminds/sprig) functions plus custom `url` and `safe` helpers. The root template object is `Template{Hermes, Email}`.
3. For HTML, generated output is CSS-inlined with [go-premailer](https://github.com/vanng822/go-premailer) unless `Hermes.DisableCSSInlining` is true.
4. `FreeMarkdown` body content (rendered via [blackfriday](https://github.com/russross/blackfriday)) replaces all body content except header/footer when set.

### Themes

A theme is any type implementing the `Theme` interface (`Name()`, `HTMLTemplate()`, `PlainTextTemplate()`). Templates are returned as **Go template strings embedded in the source** — not external files. The bundled themes:

- **`default.go`** — `Default` theme (Postmark templates).
- **`flat.go`** — `Flat` theme.
- **`cerberus.go`**, **`neopolitan.go`**, **`salted.go`** — ported from the [mailgen](https://github.com/eladnava/mailgen) themes of the same name. Mailgen ships EJS templates driven by a different data model, so these are re-implementations of each theme's *visual design* (palette/layout/fonts) on top of Hermes's own Go-template content logic — not literal translations.

`plaintext.go` holds `plainTextTemplate`, a shared plain-text template reused by the ported themes (plain-text output is theme-agnostic because it is rendered to text by `html2text`). The built-in `Default`/`Flat` themes keep their own inline copies.

To add a theme, create a new `<name>.go` implementing the interface, register it in the `testedThemes` slice at the top of `hermes_test.go` (it then inherits the full suite), and add it to the `themes` slice in `examples/main.go` so previews are generated.

### Tests

`hermes_test.go` is theme-agnostic: it runs the same suite of scenarios (`Example` interface with `getExample` / `assertHTMLContent` / `assertPlainTextContent`) against every theme in `testedThemes`. Any new theme automatically inherits the full test suite, and new common features should be added as `Example` scenarios so all themes are checked.

## Notes

- `.golangci.yml` enables nearly all linters; long-line (`lll`) checks are excluded for the template-heavy `default.go`, `flat.go`, `examples/maintenance.go`, and test files. gosec `G203` (unescaped HTML) is excluded because generating HTML is this package's purpose.
- Starting at `v1.3.0` the import path returned to `github.com/matcornic/hermes` (dropping the `/v2` suffix); `v2` tags keep the `/v2` path. See the README migration note.
