package engine_test

import (
	"github.com/exec-cmd/txtpix/internal/engine"
	"os"
	"path/filepath"
	"testing"
)

const wantBase64 = "YmFzZTY3"

func TestEncodingBase64(t *testing.T) {
	file := filepath.Join(t.TempDir(), "test_base.txt")

	if err := os.WriteFile(file, []byte("base67"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := engine.ConvertFile(file, "base64")
	if err != nil {
		t.Fatal(err)
	}

	if got != wantBase64 {
		t.Errorf("ConvertFile(file , 'base64') = %v, want %v", got, wantBase64)
	}
}
