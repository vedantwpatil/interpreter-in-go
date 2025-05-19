package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // Current pos
	readPosition int  // Current pos after reading input
	ch           byte // current char being read
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Our lexer only supports ASSCII characters instead of unicode as that would require us to introduce additional complexitiy in the way we represent and read characters
// TODO: Implement unicode and emoji support

func (l *Lexer) readChar() {
	// Signifies that we have reached the end of a file or we haven't read anything yet
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Otherwise we go to the next position in the input
		l.ch = l.input[l.readPosition]
	}
	// Moving to the next character/position
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
