package ast_test

import (
	"testing"

	"kisumu/pkg/ast"
	"kisumu/pkg/lexer"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: lexer.Token{Type: lexer.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: lexer.Token{Type: lexer.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
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
