package token

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
	)

func New(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

var keywords = map[string]TokenType{
	"fn":		FUNCTION,
	"let":		LET,
}

func LookupIdent(ident string) TokenType {
	keyword, isKeyword := keywords[ident]; if isKeyword {
		return keyword
	}
	return IDENT
}

var tokenTypes = map[byte]TokenType{
	'=':	ASSIGN,
	';':	SEMICOLON,
	'(':	LPAREN,
	')':	RPAREN,
	',':	COMMA,
	'+':	PLUS,
	'{':	LBRACE,
	'}':	RBRACE,
}

func SingleChar(ch byte) (TokenType, bool) {
	t, ok := tokenTypes[ch];
	if ok {
		return t, true
	}
	return ILLEGAL, false
}