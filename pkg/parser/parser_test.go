package parser_test

import (
	"kisumu/pkg/ast"
	"kisumu/pkg/parser"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
    let x = 5;
    let y = 10;
    let foobar = 8 * 2 + y;
    `

	l := parser.Tokenize(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	// checkParserErrors(t, l)
	if program == nil {
		t.Fatalf("Expected a program, got nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
		expectedValue      int
	}{
		{"x", 5},
		{"y", 10},
		{"foobar", 26},
	}

	for i, tc := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tc.expectedIdentifier, tc.expectedValue) {
			return
		}
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedIdentifier string, expectedValue int) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("stmt.TokenLiteral() = %q, want 'let'", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("stmt not *ast.LetStatement")
		return false
	}
	if letStmt.Name.Value != expectedIdentifier {
		t.Errorf("letStmt.Name.Value = %q, want %q", letStmt.Name.Value, expectedIdentifier)
		return false
	}
	if letStmt.Name.TokenLiteral() != expectedIdentifier {
		t.Errorf("letStmt.Name.TokenLiteral() = %q, want %q", letStmt.Name.TokenLiteral(), expectedIdentifier)
		return false
	}
	return true
}
