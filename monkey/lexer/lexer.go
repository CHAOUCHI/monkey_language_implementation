package lexer

import (
	"monkey/token"
)

/*
*
* position :
  - C'est la position du dernier character lu
  - C'est également l'index de l'attribut ch

* readPosition :
  - C'est la position du prochain character à lire
  - it can be bigger than the input when it read the entire program
*/
type Lexer struct {
	input        string
	position     int  // index of ch, the current position of the Lexer
	readPosition int  // next char position lexed
	ch           byte // current char lexed
}

/*
Instanciate a struct Lexer and read the first char
*/
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/**
* This function return a Token that match l.ch
* THEN read the next char to update ch to the next position
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
		break
	case '+':
		tok = token.New(token.PLUS, l.ch)
		break
	case '(':
		tok = token.New(token.LPAREN, l.ch)
		break
	case ')':
		tok = token.New(token.RPAREN, l.ch)
		break
	case '{':
		tok = token.New(token.LBRACE, l.ch)
		break
	case '}':
		tok = token.New(token.RBRACE, l.ch)
		break
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
		break
	case ',':
		tok = token.New(token.COMMA, l.ch)
		break
	case 0: // EOF equal 0 because the end of a string is 0
		tok.Literal = ""
		tok.Type = token.EOF
		break
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // MANDATORY early return before readChar occurs because readIdentifier already calls readChar
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // update l.ch
	return tok
}

/*
Extract a identifier or keywords from the input at position until a non letter character is hit
WARNING : You don't need to call readChar after this function because it alreay call readChar mutliple times internally
- exemple for position at 0(l) : "let five;" => "let"
- exemple for position at 4(f) : "let five;" => "five"
*/
func (l *Lexer) readIdentifier() string {
	initialPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

func (l *Lexer) readNumber() string {
	initalPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[initalPosition:l.position]
}

/**
* Read the next char from the input and increment position and readPosition
 */
func (l *Lexer) readChar() {
	if l.readPosition < len(l.input) {
		l.ch = l.input[l.readPosition]
	} else {
		l.ch = 0
	}
	// set the lexer position to the char that have been read
	l.position = l.readPosition
	// set the next readPosition to the next character
	l.readPosition += 1

	// Note :

	// I increment the readPosition in any case; even if the program have been completly read ( l.readPosition >= len(l.input))
	// otherwise the lexer could be on an infinite loop because it would be stuck at reading the last character (l.ch = l.input[l.readPosition])
	// Remember that this fonction is called by l.NextToken function wich is itself called on a while loop until EOF (0)
	// that is also why we set l.ch to 0 (EOF)
}

/*
Check if ch is a valid letter

It is used for keyword and identifier so changing the accepted character define what characters are legal for keyword and variable names on the Monkey Langage.
*/
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/*
Skip whitespace by advancing l.ch and l.readPosition until a non whitespace is encounted

whitespaces character are : SP(' '), LF('\n'), CR('\r'), HT('\t')
*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readChar()
	}
}
