package engine_test

import (
	"github.com/exec-cmd/txtpix/internal/engine"
	"testing"
)

const wantASCII = `@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@%%%#%**@==@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@##***==@**@@@...@@@....@==@@=@@=@=@@@@
@@@**===**@##@@@@.@@..@@.@@=@=@@@@@=@@@@@
@@@==#*###@%%@@@@@@@@@@@@@@=@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
`

func TestEncodingASCII(t *testing.T) {
	got, err := engine.ConvertFile("../../testdata/logo.png", "ascii")
	if err != nil {
		t.Fatal(err)
	}

	if got != wantASCII {
		t.Errorf("ConvertFile('../../testdata/logo.png' , 'ascii') = %v, want %v", got, wantASCII)
	}
}
