package tools

type (
	TokenType int
)

const (
	LeftCurl TokenType = iota
	RightCurl
	Number
	String
	COMMA
	SEMICOLUMN
	WHITESPACE
	EMPTYSTRING
	True
	False
	Null
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func NewToken(t TokenType, line int, literal string, coolumn int) Token {
	return Token{
		Type:    t,
		Literal: literal,
		Line:    line,
		Column:  coolumn,
	}
}
