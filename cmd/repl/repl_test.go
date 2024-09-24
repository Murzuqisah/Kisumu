package repl

import (
	"bytes"
	"strings"
	"testing"
)

// TestStart tests the Start function of the REPL.
func TestStart(t *testing.T) {
	input := "let x = 5;\nlet y = 10;\nlet z = x + y;\n"
	expectedOutput := []string{
		"Token: LET (let)",
		"Token: IDENT (x)",
		"Token: ASSIGN (=)",
		"Token: INT (5)",
		"Token: SEMICOLON (;)",
		"Token: LET (let)",
		"Token: IDENT (y)",
		"Token: ASSIGN (=)",
		"Token: INT (10)",
		"Token: SEMICOLON (;)",
		"Token: LET (let)",
		"Token: IDENT (z)",
		"Token: ASSIGN (=)",
		"Token: IDENT (x)",
		"Token: PLUS (+)",
		"Token: IDENT (y)",
		"Token: SEMICOLON (;)",
	}

	in := strings.NewReader(input)
	var out bytes.Buffer
	Start(in, &out)

	outputLines := strings.Split(out.String(), "\n")
	
	for i, expected := range expectedOutput {
		if i < len(outputLines) && outputLines[i] != expected {
			t.Errorf("Expected output %q, but got %q", expected, outputLines[i])
		}
	}
}
