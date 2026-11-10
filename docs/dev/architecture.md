# Architecture

txtpix is a single command, `convert`, built with [cobra](https://github.com/spf13/cobra). Module path: `github.com/exec-cmd/txtpix`.

## Layout

```
cmd/txtpix/main.go      entry point: calls cmd.Execute()
internal/cmd/          cobra root command and the convert command (flags, output)
internal/engine/       ConvertFile, the algorithm registry (Algs) and the algorithms
internal/scibe/        WriteStrInImage: writes the result text to a file
assets/logo.svg        logo used in the READMEs
testdata/logo.png      sample image used by the engine tests
```

## Conversion flow

1. `internal/cmd/convert.go` checks that `--alg` is a key of `engine.Algs`.
2. `engine.ConvertFile(path, alg)` looks up the algorithm, rejects an empty file, and opens the file.
3. The algorithm is a function `func(*os.File) (string, error)`. It returns the text.
4. `convert.go` prints the text to stdout.
5. If `--out` is set, `scibe.WriteStrInImage` writes the same text to that file.

## Where things live

- **Algorithm registry.** `Algs` in `internal/engine/engine.go` maps a name to a function. A new algorithm needs one entry there and one function in its own file.
- **Image decoding.** `ascii.go` and `unicode.go` call `image.Decode`. The blank imports of `image/png`, `image/jpeg` and `image/gif` in those files register the decoders. Add a format by importing its decoder.
- **Terminal width.** `getTerminalWidth` in `engine.go` returns the terminal width divided by 1.5, or 62 before that division when the size is unknown. That gives 41 columns for piped output.
- **Errors.** Engine errors are wrapped with `%w` and include the file name. `convert.go` returns them, and cobra prints them.

## Naming note

`internal/scibe` is named as if it writes images, but `WriteStrInImage` writes plain text.

## Tests

- `internal/engine` compares the ASCII, Unicode and base64 output of fixed inputs with golden strings.
- `internal/scibe` checks that the written file holds the text.
- `internal/cmd` has no tests yet.
