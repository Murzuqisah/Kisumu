package ast

import "kisumu/pkg/parser"

type Node interface {
	// Add methods for each node type
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
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

// statementNode method is a placeholder for any specific statement node.
type LetStatement struct {
	Token token.Token // the token.Token representing the "let" keyword
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // the token.Token representing the identifier
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type ReturnStatement struct {
	Token token.Token // the token.Token representing the "return" keyword
	Value Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
