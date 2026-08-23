# Agent Instructions for plugin-example

- Don't create binaries here — only in /tmp or ./tmp

## Language Rule

- **Chat with user:** German
- **Everything in this repository:** English (code, comments, docs, commits, tests)

## What This Is

This is the **example plugin** for the Dreego framework. It serves as a
copy-paste template for building real plugins (SSE, WebSocket, auth, etc.).

## Plugin Contract

This plugin follows the Dreego plugin contract (provisional until v1):

- Exports `Register(app *dreego.App, options Options) error`
- Uses typed Options (no `map[string]any`)
- No central Plugin interface — plugins are ordinary Go packages
- Must be called before `app.Build()` / `app.Listen()` — returns `dreego.ErrAppBuilt` otherwise
- Core never imports a plugin; the plugin imports `github.com/dreego-stack/dreego/core`

See: https://github.com/dreego-stack/dreego/blob/main/_docs/plugins.md

## Testing

- `go test ./...` — unit tests
- `go test -race ./...` — race detection
- Tests create a real `dreego.App`, register the plugin, build it, and test
  HTTP responses via `httptest.NewServer(app.Handler())`

## CI

- `.github/workflows/ci.yml` — on PR and push: `go vet`, `go test -race`, and a
  compatibility job that tests against the latest published dreego tag
- `.github/workflows/release.yml` — on push to main: validates change file,
  tests, creates changelog + tag
- `.github/dependabot.yml` — auto-updates `github.com/dreego-stack/dreego`
  dependency weekly; Dependabot PRs trigger CI, which catches incompatibility

## How to Copy This Template

1. Copy this directory: `cp -r plugin-example plugin-sse`
2. Rename the module in `go.mod`: `github.com/dreego-stack/plugin-sse`
3. Rename the Go package in `plugin.go`: `package sse`
4. Update `Options`, `Register`, and handlers for your feature
5. Update `AGENTS.md` and `README.md` with your plugin's name and purpose
6. Run `make init` to download and vendor dependencies

## Coding Rules

- Max 300 lines per handwritten file, one logical thing per file
- No code comments (except where needed for clarity)
- Go 1.22+, prefer standard library
- One Go package per repository (the plugin package at root)
- `example/` subdirectory holds a demo app, not part of the published API

## Commit Convention

Every change lands via a pull request with one `.changes/*.md` file:

```yaml
---
version: patch
---

- Feat: add X
- Bug: fix Y
```

`version: none` for changes that don't need a release tag. `version: patch`
for changes that should be released. The release workflow combines pending
files, updates the changelog, and creates a tag.