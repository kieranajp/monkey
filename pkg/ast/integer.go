package ast

import "github.com/kieranajp/monkey/pkg/token"

// IntegerLiteral is a wrapper allowing an IntegerLiteral to be used as a Statement.
type IntegerLiteral struct {
	Token token.Token // the first token of the expression.
	Value int64
}

// Implement expression
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral gets the Literal of the first token in the expression
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
