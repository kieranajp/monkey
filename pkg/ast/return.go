package ast

import (
	"bytes"

	"github.com/kieranajp/monkey/pkg/token"
)

// ReturnStatement is of the format `return <expression>;`
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// Implment the Statement interface.
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the `return` token literal.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns the contents of the ReturnStatement as a formatted string
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
