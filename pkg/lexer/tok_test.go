package lexer_test

import (
	"fmt"
	"testing"

	"kisumu/pkg/lexer"
)

type testCase struct {
	expectedTokens  lexer.TokenType
	expectedLiteral string
}

var testCases = []testCase{
	{lexer.LET, "let"},
	{lexer.IDENTIFIER, "five"},
	{lexer.ASSIGNMENT, "="},
	{lexer.INT, "5"},
	{lexer.SEMI_COLON, ";"},
	{lexer.LET, "let"},
	{lexer.IDENTIFIER, "ten"},
	{lexer.ASSIGNMENT, "="},
	{lexer.INT, "10"},
	{lexer.SEMI_COLON, ";"},
	{lexer.LET, "let"},
	{lexer.IDENTIFIER, "add"},
	{lexer.ASSIGNMENT, "="},
	{lexer.FN, "function"},
	{lexer.OPEN_PARENTHESES, "("},
	{lexer.IDENTIFIER, "x"},
	{lexer.COMMA, ","},
	{lexer.IDENTIFIER, "y"},
	{lexer.CLOSE_PARENTHESES, ")"},
	{lexer.OPEN_CURLY, "{"},
	{lexer.IDENTIFIER, "x"},
	{lexer.PLUS, "+"},
	{lexer.IDENTIFIER, "y"},
	{lexer.SEMI_COLON, ";"},
	{lexer.CLOSE_CURLY, "}"},
	{lexer.SEMI_COLON, ";"},
	{lexer.LET, "let"},
	{lexer.IDENTIFIER, "result"},
	{lexer.ASSIGNMENT, "="},
	{lexer.IDENTIFIER, "add"},
	{lexer.OPEN_PARENTHESES, "("},
	{lexer.IDENTIFIER, "five"},
	{lexer.COMMA, ","},
	{lexer.IDENTIFIER, "ten"},
	{lexer.CLOSE_PARENTHESES, ")"},
	{lexer.SEMI_COLON, ";"},
	{lexer.EOF, ""},
}

func TestGetNextTokenI(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
`
	lex := lexer.Tokenize(input)
	for i, test := range testCases {
		tok := lex.GetNextToken()
		fmt.Printf("test %d: Expected type = %s, literal = %q; got types = %s, literal = %q\n", i+1, test.expectedTokens, test.expectedLiteral, tok.Type, tok.Literal)
		if tok.Type != test.expectedTokens || tok.Literal != test.expectedLiteral {
			t.Fatalf("test[%d] - token type wrong. expected=%q (%s), got=%q (%s)",
				i, test.expectedLiteral, test.expectedTokens, tok.Literal, tok.Type)
		}
	}

}
