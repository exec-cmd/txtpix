package engine

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

var Algs = map[string]func(file *os.File) (string, error){
	"base64":  encodingBase64,
	"ascii":   encodingASCII,
	"unicode": encodingUnicode,

	"base64_decode": decodingBase64,
}

func ConvertFile(path string, alg string) (string, error) {
	fn, ok := Algs[alg]
	if !ok {
		return "", fmt.Errorf("unknown algorithm %q", alg)
	}

	isEmpty, err := fileIsEmpty(path)
	if err != nil {
		return "", fmt.Errorf("check file %q: %w", path, err)
	}
	if isEmpty {
		return "", fmt.Errorf("file %q is empty", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open file %q: %w", path, err)
	}
	defer file.Close()

	return fn(file)
}

func fileIsEmpty(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.Size() == 0, nil
}

func getTerminalWidth() int {
	terminalWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || terminalWidth < 1 {
		terminalWidth = 62
	}

	return int(float64(terminalWidth) / 1.5)
}
