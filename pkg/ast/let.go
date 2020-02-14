package ast

import "github.com/kieranajp/monkey/pkg/token"

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
