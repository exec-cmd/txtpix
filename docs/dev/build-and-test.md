# Build and test

Requires Go 1.27.1 or newer (`go.mod`).

## Makefile

| Target | What it does |
| --- | --- |
| `make fmt` | `go fmt ./...`. Rewrites files in place. |
| `make vet` | Runs `fmt`, then `go vet ./...`. |
| `make test` | Runs `vet`, then `go test ./...`. |
| `make build` | Runs `vet`, then builds `build/bin/txtpix` from `./cmd/txtpix`. This is the default target. |

`make test` and `make build` change files when `gofmt` finds something to fix. Run `git diff` after them.

## Continuous integration

`.github/workflows/ci.yml` runs on every push:

1. `gofmt -l .` must print nothing.
2. `go test ./...`
3. `go vet ./...`
4. `go build -trimpath -o /tmp/txtpix ./cmd/txtpix`

The Go version comes from `go.mod`.

## Release

`.github/workflows/release.yml` runs on tags that match `v*`:

1. Builds and tests once per target: linux/amd64, linux/arm64, windows/amd64 and darwin/arm64. Builds use `CGO_ENABLED=0`, `-trimpath` and `-ldflags="-s -w"`. The Windows binary has the `.exe` suffix.
2. Uploads each binary as a workflow artifact.
3. Creates a GitHub Release for the tag with the binaries and generated notes.

To cut a release, push a tag such as `v0.1.0`.
