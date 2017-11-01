package token

import (
	"unicode"
)

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	ILLEGAL		= "ILLEGAL"
	EOF			= "EOF"

	// identifiers + literals
	IDENT		= "IDENT"
	INT			= "INT"

	// operators
	ASSIGN		= "="
	PLUS		= "+"
	MINUS		= "-"
	BANG		= "!"
	ASTERISK	= "*"
	SLASH		= "/"
	LT			= "<"
	GT			= ">"
	LTE			= "<="
	GTE			= ">="
	EQ			= "=="
	NE			= "!="

	// delimiters
	COMMA		= ","
	SEMICOLON	= ";"

	LPAREN		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	// keywords
	FUNCTION	= "FUNCTION"
	LET			= "LET"
	TRUE		= "TRUE"
	FALSE		= "FALSE"
	IF			= "IF"
	ELSE		= "ELSE"
	RETURN		= "RETURN"
	)

func New(tokenType TokenType, ch rune) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

var operators = map[string]TokenType{
	"=":	ASSIGN,
	"+":	PLUS,
	"-":	MINUS,
	"!":	BANG,
	"*":	ASTERISK,
	"/":	SLASH,
	"<":	LT,
	"<=":	LTE,
	">":	GT,
	">=":	GTE,
	"==":	EQ,
	"!=":	NE,
}

var delimiters = map[rune]TokenType{
	'{':	LBRACE,
	'}':	RBRACE,
	';':	SEMICOLON,
	'(':	LPAREN,
	')':	RPAREN,
	',':	COMMA,
}

var keywords = map[string]TokenType{
	"fn":		FUNCTION,
	"let":		LET,
	"true":		TRUE,
	"false":	FALSE,
	"if":		IF,
	"else":		ELSE,
	"return":	RETURN,
}

func LookupIdent(ident string) TokenType {
	keyword, isKeyword := keywords[ident]; if isKeyword {
		return keyword
	}
	return IDENT
}

func LookupOperator(candidate string) TokenType {
	tt, ok := operators[candidate];
	if ok {
		return tt
	}
	return ILLEGAL
}

var operatorChar = make(map[rune]bool)

func init() {
	// populate operatorChar with char that make up all the strings for defined operators
	for str := range operators {
		for _, ch := range str {
			operatorChar[ch] = true
		}
	}
}

func IsDelimiter(ch rune) bool {
	_, ok := delimiters[ch]; if ok {
		return true
	}
	return false
}

func LookupDelimiter(ch rune) TokenType {
	tt, ok := delimiters[ch]; if ok {
		return tt
	}
	return ILLEGAL
}

func IsOperator(ch rune) bool {
	return operatorChar[ch]
}

func IsIdentifier(ch rune) bool {
	return unicode.IsLetter(ch)
}

func IsNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}