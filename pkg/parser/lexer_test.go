// test_lexer_get_next_token.go

package parser_test

// import (
// 	"fmt"
// 	"testing"

// 	"kisumu/pkg/parser"
// )

// func TestGetNextToken(t *testing.T) {
// 	input := "=+=-*/<>=!&&||...;:, \t\n\r"
// 	lexer := parser.Tokenize(input)

// 	// parser.skipWhitespace

// 	tests := []struct {
// 		expectedType parser.TokenType
// 		expectedLit  string
// 	}{
// 		{parser.ASSIGNMENT, "="},
// 		{parser.PLUS_EQUALS, "+="},
// 		{parser.DASH, "-"},
// 		{parser.STAR, "*"},
// 		{parser.SLASH, "/"},
// 		{parser.LESS, "<"},
// 		{parser.GREATER_EQUALS, ">="},
// 		{parser.NOT, "!"},
// 		{parser.AND, "&&"},
// 		{parser.OR, "||"},
// 		{parser.DOT_DOT, "..."},
// 		{parser.SEMI_COLON, ";"},
// 		{parser.COLON, ":"},
// 		{parser.COMMA, ","},
// 	}

// 	for i, test := range tests {
// 		tok := lexer.GetNextToken()
// 		fmt.Printf("test %d: Expected type = %s, literal = %q; got types = %s, literal = %q\n", i+1, test.expectedType, test, tok.Type, tok.Literal)
// 		if tok.Type != test.expectedType || tok.Literal != test.expectedLit {
// 			t.Fatalf("test[%d] - token type wrong. expected=%q (%s), got=%q (%s)",
// 				i, test.expectedLit, test.expectedType, tok.Literal, tok.Type)
// 		}
// 	}
// }
