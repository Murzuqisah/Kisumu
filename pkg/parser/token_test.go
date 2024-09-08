// test_lexer_get_next_token.go

package parser_test

import (
	"testing"

	"kisumu/pkg/parser"
)

func TestGetNextToken(t *testing.T) {
	input := "=+=-*/<>=!&|...;:, \t\n\r"
	lexer := parser.Tokenize(input)

	tests := []struct {
		expectedType parser.TokenType
		expectedLit  string
	}{
		{parser.ASSIGNMENT, "="},
		// {parser.PLUS, "+"},
		{parser.PLUS_EQUALS, "+="},
		{parser.DASH, "-"},
		// {parser.MINUS_MINUS, "--"},
		{parser.STAR, "*"},
		{parser.SLASH, "/"},
		// {parser.SLASH_EQUALS, "/="},
		// {parser.STAR_EQUALS, "*="},
		{parser.LESS, "<"},
		// {parser.LESS_EQUAL, "<="},
		// {parser.GREATER, ">"},
		{parser.GREATER_EQUALS, ">="},
		// {parser.NOT_EQUALS, "!="},
		{parser.NOT, "!"},
		{parser.AND, "&"},
		{parser.OR, "|"},
		{parser.DOT_DOT, "..."},
		{parser.SEMI_COLON, ";"},
		{parser.COLON, ":"},
		// {parser.QUESTION, "?"},
		{parser.COMMA, ","},
		{parser.WHITESPACE, " "},
		{parser.WHITESPACE, "\t"},
		{parser.WHITESPACE, "\n"},
		{parser.WHITESPACE, "\r"},
		{parser.EOF, ""},
	}

	for i, test := range tests {
		tok := lexer.GetNextToken()
		if tok.Type != test.expectedType || tok.Literal != test.expectedLit {
			t.Fatalf("test[%d] - token type wrong. expected=%q (%s), got=%q (%s)",
				i, test.expectedLit, test.expectedType, tok.Literal, tok.Type)
		}
	}
}
