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
		"kisumu $Token: LET (let)",
		"Token: IDENTIFIER (x)",
		"Token: ASSIGNMENT (=)",
		"Token: INT (5)",
		"Token: SEMI_COLON (;)",
		"kisumu $Token: LET (let)",
		"Token: IDENTIFIER (y)",
		"Token: ASSIGNMENT (=)",
		"Token: INT (10)",
		"Token: SEMI_COLON (;)",
		"kisumu $Token: LET (let)",
		"Token: IDENTIFIER (z)",
		"Token: ASSIGNMENT (=)",
		"Token: IDENTIFIER (x)",
		"Token: PLUS (+)",
		"Token: IDENTIFIER (y)",
		"Token: SEMI_COLON (;)",
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
