package token

// Type identifies the type of token we're working with: number, var name, fn, etc
type Type string

// Token represents a piece of code parsed by the lexer.
type Token struct {
	Type    Type
	Literal string
}

// Enum of all possible token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers & literals
	IDENT = "IDENT" // add, foobar, x, y...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks if an identifier is a keyword or not
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
