English · [Русский](README.RU.md)

<p align="center">
    <img src="assets/logo.svg" width="360" alt="logo txtpix">
</p>

# txtpix

Turn images into text art and files into base64, from the terminal.

[![CI](https://github.com/exec-cmd/txtpix/actions/workflows/ci.yml/badge.svg)](https://github.com/exec-cmd/txtpix/actions/workflows/ci.yml)
[![Version Go](https://img.shields.io/github/go-mod/go-version/exec-cmd/txtpix)](go.mod)
[![License: MIT](https://img.shields.io/github/license/exec-cmd/txtpix)](LICENSE)
[![Last release](https://img.shields.io/github/v/release/exec-cmd/txtpix)](https://github.com/exec-cmd/txtpix/releases)

## Features

- **ASCII art.** Maps image brightness to a ramp of characters: `@%#*+=-:. `.
- **Unicode art.** Packs two pixel rows into one line with half-block characters (`▀ ▄ █`).
- **base64.** Encodes any file to base64 text, and decodes it back.
- **Output.** Prints to stdout, or writes the same text to a file with `-o`.

## Install

Requires Go 1.27.1 or newer.

### Build from source

```sh
make build
./build/bin/txtpix --help
```

Without `make`:

```sh
go build -o txtpix ./cmd/txtpix
```

### go install

```sh
go install github.com/exec-cmd/txtpix/cmd/txtpix@latest
```

### Release binaries

Each `v*` tag publishes binaries on the GitHub Releases page for linux/amd64, linux/arm64, windows/amd64 and darwin/arm64.

## Usage

```sh
txtpix convert <file> [flags]
```

| Flag | Default | Description |
| --- | --- | --- |
| `-a`, `--alg` | `base64` | Algorithm: `base64`, `ascii`, `unicode`, `base64_decode` |
| `-o`, `--out` | none | Also write the result to this file |

### ASCII art

```sh
txtpix convert testdata/logo.png -a ascii
```

```text
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@%%%#%**@==@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@##***==@**@@@...@@@....@==@@=@@=@=@@@@
@@@**===**@##@@@@.@@..@@.@@=@=@@@@@=@@@@@
@@@==#*###@%%@@@@@@@@@@@@@=@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
```

### Unicode art

```sh
txtpix convert testdata/logo.png -a unicode
```

```text
█▀                                      ▀
           ▄▄                            
        ▄▄ ▀▀                            
        ▀▀      ▀█▀ ▄▄▀▀█▀ █▀▄ ▀  ▀▄▀    
     ███         █▄ ▀▀▄ █▄ █▄▀ ▄▄ ▄▀▄    
   ██                      ▀             
▄                                        
```

### base64 and decoding

```sh
printf 'hello' > hello.txt
txtpix convert hello.txt -o hello.b64   # prints aGVsbG8= and writes it to hello.b64
txtpix convert -a base64_decode hello.b64
```

### Save the result

```sh
txtpix convert testdata/logo.png -a ascii -o logo.txt
```

`-o` writes the text to the file. It does not create an image.

## Notes

- **Width.** Image algorithms use the terminal width divided by 1.5. When the width is unknown, for example when output is piped, they use 41 columns.
- **Image formats.** `ascii` and `unicode` read PNG, JPEG and GIF.
- **Empty files** are rejected with an error.
- **`base64_decode`** prints the decoded bytes as text, so binary data comes out garbled.

## Development

```sh
make test     # format, vet and run the tests
make build    # build build/bin/txtpix
```

## License

MIT. See [LICENSE](LICENSE).
