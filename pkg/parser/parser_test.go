package parser_test

import (
	"testing"

	"kisumu/pkg/parser"
)

type Token struct {
	expectedType    parser.TokenType
	expectedLiteral string
}

var expectedTokens = []Token{
	{
		expectedType:    parser.OPEN_BRACKET,
		expectedLiteral: "[",
	},
	{
		expectedType:    parser.IDENTIFIER,
		expectedLiteral: "hello",
	},
	{
		expectedType:    parser.CLOSE_BRACKET,
		expectedLiteral: "]",
	},
}

func TestTokenize(t *testing.T) {
	source := `[hello ]` // l.skipWhitespace() is defined in parser/lexer.go, it skips whitespace before parsing tokens

	ksm := parser.Tokenize(source)

	for i, tt := range expectedTokens {
		tokens := ksm.GetNextToken()
		if string(tokens.Type) != string(tt.expectedType) || tokens.Literal != tt.expectedLiteral {
			t.Fatalf("Expected token %d to be {%v, %s}, but got {%v, %s}", i, tt.expectedType, tt.expectedLiteral, tokens.Type, tokens.Literal)
		}
	}
}
