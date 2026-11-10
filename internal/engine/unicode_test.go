package engine_test

import (
	"github.com/exec-cmd/txtpix/internal/engine"
	"testing"
)

const wantUnicode = `█▀                                      ▀
           ▄▄                            
        ▄▄ ▀▀                            
        ▀▀      ▀█▀ ▄▄▀▀█▀ █▀▄ ▀  ▀▄▀    
     ███         █▄ ▀▀▄ █▄ █▄▀ ▄▄ ▄▀▄    
   ██                      ▀             
▄                                        
`

func TestEncodingUnicode(t *testing.T) {
	got, err := engine.ConvertFile("../../testdata/logo.png", "unicode")
	if err != nil {
		t.Fatal(err)
	}

	if got != wantUnicode {
		t.Errorf("ConvertFile('../../testdata/logo.png' , 'unicode') = %v, want %v", got, wantUnicode)
	}
}
