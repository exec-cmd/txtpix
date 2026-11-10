package engine

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"
)

func encodingASCII(file *os.File) (string, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("decode image %q: %w", file.Name(), err)
	}

	const chars = "@%#*+=-:. "

	bounds := img.Bounds()
	width := min(bounds.Dx(), getTerminalWidth())

	height := max(1, bounds.Dy()*width/bounds.Dx()/2)

	var result strings.Builder

	for y := range height {
		for x := range width {
			sourceX := bounds.Min.X + x*bounds.Dx()/width
			sourceY := bounds.Min.Y + y*bounds.Dy()/height

			r, g, b, _ := img.At(sourceX, sourceY).RGBA()
			brightness := (299*r + 587*g + 114*b) / 1000
			index := brightness * uint32(len(chars)-1) / 65535

			result.WriteByte(chars[index])
		}
		result.WriteString("\n")
	}

	return result.String(), nil
}
