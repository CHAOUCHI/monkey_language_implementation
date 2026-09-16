package lexer

import "monkey/token"

/**
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

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/**
* This function return a Token that match l.ch
* THEN update l.ch
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = token.New(token.ASSIGN, l.ch)
	case '+':
		tok = token.New(token.PLUS, l.ch)
	case '(':
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		tok = token.New(token.RPAREN, l.ch)
	case '{':
		tok = token.New(token.LBRACE, l.ch)
	case '}':
		tok = token.New(token.RBRACE, l.ch)
	case ';':
		tok = token.New(token.SEMICOLON, l.ch)
	case 0: // EOF equal 0 because the end of a string is 0
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar() // update l.ch
	return tok
}

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
