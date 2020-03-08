package ast

import (
	"bytes"

	"github.com/kieranajp/monkey/pkg/token"
)

// LetStatement is a Node representation of a variable assignment.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// This is just here because Go's interfaces can eat a dick
func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal field of the LET token.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String prints out the Let statement as a formatted string.
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
