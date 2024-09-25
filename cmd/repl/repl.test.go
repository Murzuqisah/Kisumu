package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestStartI(t *testing.T) {
	input := "let x = 5;\n"
	expectedOutput := "Token: LET (let)\nToken: IDENTIFIER (x)\nToken: ASSIGNMENT (=)\nToken: INT (5)\nToken: SEMI_COLON (;)\n"

	in := strings.NewReader(input)
	var out bytes.Buffer
	Start(in, &out)

	if out.String() != expectedOutput {
		t.Errorf("Expected output %q, but got %q", expectedOutput, out.String())
	}
}
