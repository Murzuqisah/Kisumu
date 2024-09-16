package lexer_test

import (
	"testing"

	"kisumu/pkg/lexer"
)

type Token struct {
	expectedType    lexer.TokenType
	expectedLiteral string
}

var expectedTokens = []Token{
	{
		expectedType:    lexer.OPEN_BRACKET,
		expectedLiteral: "[",
	},
	{
		expectedType:    lexer.IDENTIFIER,
		expectedLiteral: "hello",
	},
	{
		expectedType:    lexer.CLOSE_BRACKET,
		expectedLiteral: "]",
	},
}

func TestTokenize(t *testing.T) {
	source := `[hello ]` // l.skipWhitespace() is defined in lexer/lexer.go, it skips whitespace before parsing tokens

	ksm := lexer.Tokenize(source)

	for i, tt := range expectedTokens {
		tokens := ksm.GetNextToken()
		if string(tokens.Type) != string(tt.expectedType) || tokens.Literal != tt.expectedLiteral {
			t.Fatalf("Expected token %d to be {%v, %s}, but got {%v, %s}", i, tt.expectedType, tt.expectedLiteral, tokens.Type, tokens.Literal)
		}
	}
}
