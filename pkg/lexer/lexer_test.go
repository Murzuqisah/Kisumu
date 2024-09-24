package lexer_test

import (
	"fmt"
	"testing"

	"kisumu/pkg/lexer"
)

func TestGetNextToken(t *testing.T) {
	input := "=+=-*/<>=!&&||...;:, \t\n\r"
	lex := lexer.Tokenize(input)

	tests := []struct {
		expectedType lexer.TokenType
		expectedLit  string
	}{
		{lexer.ASSIGNMENT, "="},
		{lexer.PLUS_EQUALS, "+="},
		{lexer.DASH, "-"},
		{lexer.ASTERISK, "*"},
		{lexer.SLASH, "/"},
		{lexer.LESS, "<"},
		{lexer.GREATER_EQUALS, ">="},
		{lexer.BANG, "!"},
		{lexer.AND, "&&"},
		{lexer.OR, "||"},
		{lexer.DOT_DOT, "..."},
		{lexer.SEMI_COLON, ";"},
		{lexer.COLON, ":"},
		{lexer.COMMA, ","},
	}

	for i, test := range tests {
		tok := lex.GetNextToken()
		fmt.Printf("test %d: Expected type = %s, literal = %q; got types = %s, literal = %q\n", i+1, test.expectedType, test.expectedLit, tok.Type, tok.Literal)
		if tok.Type != test.expectedType || tok.Literal != test.expectedLit {
			t.Fatalf("test[%d] - token type wrong. expected=%q (%s), got=%q (%s)",
				i, test.expectedLit, test.expectedType, tok.Literal, tok.Type)
		}
	}
}
