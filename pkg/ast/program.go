package ast

// Program is the root node of the AST
type Program struct {
	Statements []Statement
}

// TokenLiteral is mostly used for debugging purposes, to read the literal value of the Node's token.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
