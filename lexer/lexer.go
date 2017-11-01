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

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	ch, _ := utf8.DecodeRuneInString(l.input[l.readPosition:])
	return ch
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhiteSpace()
	if l.ch == 0 {
		return token.Token{Type: token.EOF, Literal:""}
	}
	if token.IsOperator(l.ch) {
		tok.Literal = l.readOperator()
		tok.Type = token.LookupOperator(tok.Literal)
	} else if token.IsIdentifier(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
	} else if token.IsNumber(l.ch) {
		tok.Type = token.INT
		tok.Literal = l.readNumber()
	} else if token.IsDelimiter(l.ch) {
		tt := token.LookupDelimiter(l.ch)
		tok = token.New(tt, l.ch)
		l.readChar() // advance past delimiter
	} else {
		tok = token.New(token.ILLEGAL, l.ch)
		l.readChar() // skip past illegal token
	}
	return
}

func (l *Lexer) readOperator() string {
	startPos := l.position
	for token.IsOperator(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for (unicode.IsSpace(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for token.IsNumber(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for token.IsIdentifier(l.ch) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}