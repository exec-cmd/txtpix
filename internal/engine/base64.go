package engine

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func encodingBase64(file *os.File) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read file %q: %w", file.Name(), err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func decodingBase64(file *os.File) (string, error) {
	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("read file %q: %w", file.Name(), err)
	}

	out, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return "", fmt.Errorf("decode base64 %q: %w", file.Name(), err)
	}
	return string(out), nil
}
