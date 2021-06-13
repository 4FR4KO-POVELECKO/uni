package compiler

type Token int

const (
	EOF = iota
	ILLEGAL
	VAR
	INT
	SEMI // ;

	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =
	EXC    // !
	DOT    // .
	COMMA  // ,
	LPAR   // (
	RPAR   // )
	LCUR   // {
	RCUR   // }
	LSQU   // [
	RSQU   // ]
	QUOS   // "
	APOS   // '
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	VAR:     "VAR",
	INT:     "INT",
	SEMI:    ";",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
	EXC:    "!",
	DOT:    ".",
	COMMA:  ",",
	LPAR:   "(",
	RPAR:   ")",
	LCUR:   "{",
	RCUR:   "}",
	LSQU:   "[",
	RSQU:   "]",
	QUOS:   "\"",
	APOS:   "'",
}

func (t Token) String() string {
	return tokens[t]
}
