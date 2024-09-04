package test_test

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
		expectedType:    parser.STRING,
		expectedLiteral: "hello",
	},
	{
		expectedType:    parser.WHITESPACE,
		expectedLiteral: " ",
	},
	{
		expectedType:    parser.CLOSE_BRACKET,
		expectedLiteral: "]",
	},
}

func TestTokenize(t *testing.T) {
	source := "[hello] world"

	ksm := parser.Tokenize(source)
	tokens := ksm

	if len(tokens) != len(expectedTokens) {
		t.Fatalf("Expected %d tokens, but got %d", len(expectedTokens), len(tokens))
	}

	for i, tt := range expectedTokens {
		if string(tokens[i].Type) != string(tt.expectedType) || tokens[i].Literal != tt.expectedLiteral {
			t.Fatalf("Expected token %d to be {%v, %s}, but got {%v, %s}", i, tt.expectedType, tt.expectedLiteral, tokens[i].Type, tokens[i].Literal)
		}
	}
}
