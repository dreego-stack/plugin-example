# plugin-example

Example plugin for the [Dreego](https://github.com/dreego-stack/dreego) framework.
Copy this repository to start building your own Dreego plugin.

## What It Does

Registers a greeting endpoint, a health check, a custom 404 handler, and
optional request logging on a Dreego app. The code is intentionally simple —
it demonstrates the plugin contract, not a real feature.

## Usage

```go
package main

import (
    "log"
    "os"

    dreego "github.com/dreego-stack/dreego/core"
    example "github.com/dreego-stack/plugin-example"
)

func main() {
    app := dreego.New()
    if err := example.Register(app, example.Options{
        Prefix:        "/api",
        Greeting:      "Hi",
        EnableLogging: true,
    }); err != nil {
        log.Fatal(err)
    }
    if err := app.Listen(":8080"); err != nil {
        log.Fatal(err)
    }
}
```

Then:

```
curl http://localhost:8080/api/greet/World
# Hi, World!

curl http://localhost:8080/api/health
# ok
```

## Options

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Prefix` | `string` | `"/example"` | Route prefix for all plugin routes |
| `Greeting` | `string` | `"Hello"` | Greeting word used by the greet endpoint |
| `EnableLogging` | `bool` | `false` | Enable Dreego request-logging middleware |

## Getting Started

```sh
make init    # download and vendor dependencies (one-time)
make test    # run tests using vendored deps (offline)
make run     # run the demo app
```

The plugin resolves `github.com/dreego-stack/dreego` from the Go module proxy
(published tags), just like any external consumer. After `make init`, all
dependencies are vendored in `vendor/` and the plugin works offline — no
local dreego checkout needed.

## CI

- **ci.yml** — runs `go vet`, `go test -race`, and a compatibility check against
  the latest published Dreego tag on every PR and push.
- **release.yml** — validates the `.changes/*.md` file, runs tests, creates a
  changelog commit and a git tag on push to main.
- **dependabot.yml** — opens PRs to update the `github.com/dreego-stack/dreego`
  dependency weekly. If a Dreego release breaks this plugin, the Dependabot PR's
  CI run fails and the maintainer is notified.

## License

MPL-2.0, same as Dreego.