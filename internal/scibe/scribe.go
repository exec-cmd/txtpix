package scibe

import (
	"os"
)

func WriteStrInImage(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0o644)
	if err != nil {
		return err
	}
	return nil
}
