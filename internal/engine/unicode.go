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

func encodingUnicode(file *os.File) (string, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("decode image %q: %w", file.Name(), err)
	}

	bounds := img.Bounds()
	width := min(bounds.Dx(), getTerminalWidth())
	height := max(1, bounds.Dy()*width/bounds.Dx()/2)

	var result strings.Builder

	for y := range height {
		for x := range width {
			sourceX := bounds.Min.X + x*bounds.Dx()/width
			topY := bounds.Min.Y + (2*y)*bounds.Dy()/(2*height)
			bottomY := bounds.Min.Y + (2*y+1)*bounds.Dy()/(2*height)

			topR, topG, topB, topA := img.At(sourceX, topY).RGBA()
			bottomR, bottomG, bottomB, bottomA := img.At(sourceX, bottomY).RGBA()

			topR += 0xffff - topA
			topG += 0xffff - topA
			topB += 0xffff - topA

			bottomR += 0xffff - bottomA
			bottomG += 0xffff - bottomA
			bottomB += 0xffff - bottomA

			lum := func(r, g, b uint32) uint32 { return (299*r + 587*g + 114*b) / 1000 }
			top := lum(topR, topG, topB) > 0x8000
			bottom := lum(bottomR, bottomG, bottomB) > 0x8000

			var ch string
			switch {
			case top && bottom:
				ch = "█"
			case top:
				ch = "▀"
			case bottom:
				ch = "▄"
			default:
				ch = " "
			}
			result.WriteString(ch)

		}
		result.WriteString("\n")
	}

	return result.String(), nil

}
