package ast

import (
	"github.com/thebashshell/simply/token"
)

type Node interface {
	TokenValue() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenValue() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenValue()
	}

	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) TokenValue() string {
	return i.Token.Value
}

func (i *Identifier) expressionNode() {}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenValue() string {
	return l.Token.Value
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()     {}
func (rs *ReturnStatement) TokenValue() string { return rs.Token.Value }

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
