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
	EQ       = "=="
	NOT_EQ   = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// List of the keywords of the MonkeyLangage
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tokenType, isKeyword := keywords[ident]; isKeyword {
		return tokenType
	}
	return IDENT
}
