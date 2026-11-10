# Algorithms

Select an algorithm with `-a` / `--alg`. The default is `base64`. The names are the keys of `engine.Algs` in `internal/engine/engine.go`.

| Name | Input | Output | Source file |
| --- | --- | --- | --- |
| `base64` | any file | standard base64 text (RFC 4648, with padding) | `base64.go` |
| `base64_decode` | base64 text | the decoded bytes, as a string | `base64.go` |
| `ascii` | PNG, JPEG or GIF | ASCII art, one character per cell | `ascii.go` |
| `unicode` | PNG, JPEG or GIF | Unicode half-block art, two pixel rows per line | `unicode.go` |

## base64

Reads the whole file and encodes it with `base64.StdEncoding`.

## base64_decode

Decodes the whole file with `base64.StdEncoding` and returns the bytes as a Go string. The output is printed with `fmt.Println`, so binary data is not restored. Only text payloads round-trip. Line breaks are skipped. Other invalid input returns an error.

## ascii

- Size: the width is the smaller of the image width and the output width from `getTerminalWidth`. The height is `width × image height / image width / 2`, because a terminal cell is about twice as tall as it is wide.
- Sampling: nearest-neighbour. Each output cell reads one pixel.
- Ramp: `@%#*+=-:. ` from dark to light. Brightness is `(299R + 587G + 114B) / 1000`, computed on the 16-bit colour values from `RGBA()`.

## unicode

- Size: the same width rule as `ascii`. Each output line covers two rows of pixels.
- Transparency: alpha is composited over white before the brightness test.
- Cells: a top and bottom pixel that are both bright give `█`, only the top gives `▀`, only the bottom gives `▄`, and neither gives a space. Bright means luminance above `0x8000`.

## Errors

| Case | Message |
| --- | --- |
| unknown algorithm | `unknown algorithm "<name>"` |
| empty file | `file "<path>" is empty` |
| file cannot be stat'ed or opened | `check file` / `open file` with the path |
| undecodable image | `decode image "<path>"` |
| invalid base64 | `decode base64 "<path>"` |
