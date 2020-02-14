package ast

// Node must be implemented by every node in the AST.
type Node interface {
	TokenLiteral() string
}

// A Statement is a Node that does not produce a value
type Statement interface {
	Node
	statementNode()
}

// An Expression is a Node that produces a value
type Expression interface {
	Node
	expressionNode()
}
