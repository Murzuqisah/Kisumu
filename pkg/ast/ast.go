package ast

import (
	"bytes"
	"strings"

	"kisumu/pkg/lexer"
)

type PrefixExpression struct {
	Token    lexer.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    lexer.TokenType // the token.Token representing the operator
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

func (oe *InfixExpression) TokenLiteral() string {
	return oe.Right.TokenLiteral()
}

func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type IntegerLiteral struct {
	Token lexer.Token // the token.Token representing the integer value
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// Program struct represents the entire program.
type Program struct {
	Statements []Statement
}

// TokenLiteral method returns the literal value of the token.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

// statementNode method is a placeholder for any specific statement node.
type LetStatement struct {
	Token lexer.Token // the token.Token representing the "let" keyword
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token lexer.Token // the token.Token representing the identifier
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       lexer.Token // the token.Token representing the "return" keyword
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

type IfExpression struct {
	Token       lexer.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type BlockStatement struct {
	Token      lexer.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      lexer.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	param := []string{}
	for _, p := range fl.Parameters {
		param = append(param, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(param, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token    lexer.Token
	Function Expression
	Argument []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Argument {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
