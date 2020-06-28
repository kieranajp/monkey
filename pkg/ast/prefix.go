package ast

import (
	"bytes"

	"github.com/kieranajp/monkey/pkg/token"
)

// PrefixExpression represents a prefixed Expression, e.g. !true or -15
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral gets the Prefix TokenLiteral
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString("}")

	return out.String()
}
