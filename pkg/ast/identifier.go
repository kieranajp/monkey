package ast

import "github.com/kieranajp/monkey/pkg/token"

// An Identifier is the AST Node representation of an IDENT token.
type Identifier struct {
	Token token.Token
	Value string
}

// This is just here because Go's interfaces can eat a dick.
func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal field of the IDENT token.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String up the token value
func (i *Identifier) String() string { return i.Value }
