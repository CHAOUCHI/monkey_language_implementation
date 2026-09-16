package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "=+(){};"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, it := range tests {

		var tok token.Token = l.NextToken()
		if it.expectedType != tok.Type {
			t.Fatalf("tests[%d] - tokenType wrong. expected %q got %q", i, it.expectedType, tok.Type)
		}

		if it.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d] - tokenLiteral wrong. expected %q got %q", i, it.expectedLiteral, tok.Literal)
		}
	}
}
