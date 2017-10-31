package lexer

import (
	"theduke/token"
)

type Lexer struct {
	input			string
	position		int		// current position in input
	readPosition	int		// current reading position in input
	ch				byte	// current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}


func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()
	if l.ch == 0 {
		return token.Token{Type: token.EOF, Literal:""}
	}
	t, isSingle := token.SingleChar(l.ch); if isSingle {
		tok = token.New(t, l.ch)
	} else if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return // readIdentifier has already moved us to next position
	} else if isDigit(l.ch) {
		tok.Type = token.INT
		tok.Literal = l.readNumber()
		return // readNumber has already moved us to the next position
	} else {
		tok = token.New(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}