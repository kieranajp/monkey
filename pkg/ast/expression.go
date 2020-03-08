package ast

import "github.com/kieranajp/monkey/pkg/token"

// ExpressionStatement is a wrapper allowing an Expression to be used as a Statement.
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression.
	Expression Expression
}

// Implement statement
func (es *ExpressionStatement) statementNode() {}

// TokenLiteral gets the Literal of the first token in the expression
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
