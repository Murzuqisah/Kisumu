package ast_test

import (
	"testing"

	"kisumu/pkg/lexer"
)

// let myVar = anotherVar;

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: lexer.Token{Type: lexer.LET, Literal: "let"},
				Name: &Identifier{
					Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
					51},
				Value: &Identifier{
					Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
