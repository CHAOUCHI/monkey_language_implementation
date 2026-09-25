package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, tokenLiteral byte) Token {
	return Token{tokenType, string(tokenLiteral)}
}

// List of the TokenTypes of the MonkeyLangage
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + variable type
	IDENT = "IDENT" // the variable names
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

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

// List of the keywords of the MonkeyLangage
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tokenType, isKeyword := keywords[ident]; isKeyword {
		return tokenType
	}
	return IDENT
}
