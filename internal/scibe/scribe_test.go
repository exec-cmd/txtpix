package scibe_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/exec-cmd/txtpix/internal/scibe"
)

func TestWriteStrInImage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test_base.txt")
	err := scibe.WriteStrInImage(path, "67")
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.ReadFile(path)

	if err != nil {
		t.Fatal(err)
	}

	if string(file) != "67" {
		t.Fatal(fmt.Errorf("TestWriteStrInImage invalid working"))
	}

}
