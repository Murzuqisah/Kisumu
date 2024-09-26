package parser_test

import (
	"testing"

	"kisumu/pkg/ast"
	"kisumu/pkg/lexer"
)

func TestBangToken(t *testing.T) {
	input := "!5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
	}

	if exp.Operator != "!" {
		t.Fatalf("exp.Operator is not '!'. got=%s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestDashTokenHandling(t *testing.T) {
	input := "-5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
	}

	if exp.Operator != "-" {
		t.Fatalf("exp.Operator is not '-'. got=%s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestParseIdentifier(t *testing.T) {
	input := "identifier;"

	l := lexer.Tokenize(input)
	p := NewParser(l)

	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. Got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. Got %T", stmt.Expression)
	}

	if ident.Value != "identifier" {
		t.Errorf("ident.Value not %s. Got %s", "identifier", ident.Value)
	}

	if ident.TokenLiteral() != "identifier" {
		t.Errorf("ident.TokenLiteral not %s. Got %s", "identifier", ident.TokenLiteral())
	}
}

func TestParseIntegerLiteral(t *testing.T) {
	input := "5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)

	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func TestParseInfixExpressionPlus(t *testing.T) {
	input := "5 + 5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. Got %d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not ast.InfixExpression. Got %T", stmt.Expression)
	}

	if !testIntegerLiteral(t, exp.Left, 5) {
		return
	}

	if exp.Operator != "+" {
		t.Fatalf("exp.Operator is not '+'. Got %s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestSlashToken(t *testing.T) {
	input := "5 / 2;"
	l := lexer.Tokenize(input)
	p := NewParser(l)

	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. Got %d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not ast.InfixExpression. Got %T", stmt.Expression)
	}

	if exp.Operator != "/" {
		t.Fatalf("exp.Operator is not '/'. Got %s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Left, 5) {
		return
	}

	if !testIntegerLiteral(t, exp.Right, 2) {
		return
	}
}

func TestParseAsteriskToken(t *testing.T) {
	input := "*5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
	}

	if exp.Operator != "*" {
		t.Fatalf("exp.Operator is not '*'. got=%s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestParseInfixExpressionEquals(t *testing.T) {
	input := "5 == 5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. Got %d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not ast.InfixExpression. Got %T", stmt.Expression)
	}

	if !testIntegerLiteral(t, exp.Left, 5) {
		return
	}

	if exp.Operator != "==" {
		t.Fatalf("exp.Operator is not '=='. Got %s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestNotEqualsInfixExpression(t *testing.T) {
	input := "5 != 5;"

	l := lexer.Tokenize(input)
	p := NewParser(l)
	program := p.ParseProgram()
	CheckParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. Got %d\n", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not ast.InfixExpression. Got %T", stmt.Expression)
	}

	if !testIntegerLiteral(t, exp.Left, 5) {
		return
	}

	if exp.Operator != "!=" {
		t.Fatalf("exp.Operator is not '!='. Got %s", exp.Operator)
	}

	if !testIntegerLiteral(t, exp.Right, 5) {
		return
	}
}

func TestParsingInfixExpressionsWithLessGreater(t *testing.T) {
	infixTest := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 > 3;", 5, ">", 3},
		{"3 < 5;", 3, "<", 5},
	}

	for _, tt := range infixTest {
		l := lexer.Tokenize(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		CheckParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. Got %d\n", 1, len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. Got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. Got %T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. Got %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

// func CheckParserErrors(t *testing.T, p *Parser) {
// 	errors := p.Errors()

// 	if len(errors) == 0 {
// 		return
// 	}

// 	t.Errorf("parser has %d errors", len(errors))
// 	for _, msg := range errors {
// 		t.Errorf("parser error: %q", msg)
// 	}
// 	t.FailNow()
// }

// func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
// 	integ, ok := il.(*ast.IntegerLiteral)
// 	if !ok {
// 		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
// 		return false
// 	}
// 	if integ.Value != value {
// 		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
// 		return false
// 	}
// 	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
// 		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
// 			integ.TokenLiteral())
// 		return false
// 	}
// 	return true
// }
