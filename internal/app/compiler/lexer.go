package compiler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

type Position struct {
	line   int
	column int
}

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 0},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) LexerScan() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}

			panic(err)
		}

		l.pos.column++

		switch r {
		case '\n':
			l.resetPosition()
		case ';':
			return l.pos, SEMI, ";"
		case '+':
			return l.pos, ADD, "+"
		case '-':
			return l.pos, SUB, "-"
		case '*':
			return l.pos, MUL, "*"
		case '/':
			return l.pos, DIV, "/"
		case '=':
			return l.pos, ASSIGN, "="
		case '!':
			return l.pos, EXC, "!"
		case '?':
			return l.pos, QUE, "?"
		case '.':
			return l.pos, DOT, "."
		case ',':
			return l.pos, COMMA, ","
		case '_':
			return l.pos, UNSCORE, "_"
		case '(':
			return l.pos, LPAR, "("
		case ')':
			return l.pos, RPAR, ")"
		case '[':
			return l.pos, LSQU, "["
		case ']':
			return l.pos, RSQU, "]"
		case '{':
			return l.pos, LCUR, "{"
		case '}':
			return l.pos, RCUR, "}"
		case '<':
			return l.pos, LQUO, "<"
		case '>':
			return l.pos, RQUO, ">"
		case '"':
			return l.pos, QUOS, "\""
		case '\'':
			return l.pos, APOS, "'"
		default:
			if unicode.IsSpace(r) {
				continue
			} else if unicode.IsDigit(r) {
				startPosition := l.pos
				l.backup()
				lit := l.lexInt()

				return startPosition, INT, lit
			} else if unicode.IsLetter(r) {
				startPosition := l.pos
				l.backup()
				lit := l.lexVar()

				return startPosition, VAR, lit
			} else {
				return l.pos, ILLEGAL, string(r)
			}
		}
	}
}

func (l *Lexer) resetPosition() {
	l.pos.line++
	l.pos.column = 0
}

func (l *Lexer) backup() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.pos.column--
}

func (l *Lexer) lexInt() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++

		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else if r == '.' {
			lit = lit + string(r)
		} else if r == ',' {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}

	}
}

func (l *Lexer) lexVar() string {
	var lit string

	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.column++

		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else if r == UNSCORE {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}

	}
}

func Main() {
	fmt.Println("Uni!")

	file, err := os.Open("test.uni")
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.LexerScan()
		if tok == EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	}
}
