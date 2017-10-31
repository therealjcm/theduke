package lexer

import (
	"theduke/token"
	"unicode/utf8"
	"unicode"
)

type Lexer struct {
	input			string
	position		int		// current position in input
	readPosition	int		// current reading position in input
	ch				rune	// current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	var width int
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch, width = utf8.DecodeRuneInString(l.input[l.readPosition:])
	}
	l.position = l.readPosition
	l.readPosition += width
}


func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()
	if l.ch == 0 {
		return token.Token{Type: token.EOF, Literal:""}
	}
	t, isSingle := token.SingleChar(l.ch); if isSingle {
		tok = token.New(t, l.ch)
	} else if unicode.IsLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return // readIdentifier has already moved us to next position
	} else if unicode.IsDigit(l.ch) {
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

func (l *Lexer) readNumber() string {
	startPos := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}