package eval

import (
	"errors"
	"strconv"
	"strings"
)

type lexer struct {
	data []rune
	pos  int
	len  int
	res  []Token
	ctx  *Context
}

func newLexer(data []rune, ctx *Context) *lexer {
	return &lexer{data: data, pos: 0, len: len(data), res: []Token{}, ctx: ctx}
}

// lex goes through data and sets res to result
func (l *lexer) lex() error {
	for l.pos = 0; l.pos < l.len; l.pos++ {
		err := l.next()
		if err != nil {
			return err
		}
	}
	return nil
}

// next gets the next token from the data
func (l *lexer) next() error {
	cur := l.data[l.pos]
	var err error
	switch {
	case IsLeft(cur):
		l.add(Token{LEFT, nil})
	case IsRight(cur):
		l.add(Token{RIGHT, nil})
	case cur == '+':
		l.add(Token{NUMOP, ADD})
	case cur == '-':
		l.add(Token{NUMOP, SUB})
	case cur == '*':
		l.add(Token{NUMOP, MUL})
	case cur == '/':
		l.add(Token{NUMOP, DIV})
	case cur == '^':
		l.add(Token{NUMOP, POW})
	case cur == '%':
		l.add(Token{NUMOP, MOD})
	case cur == '=':
		l.add(Token{COMP, EQUAL})
	case cur == '!':
		l.handleDoubleSingle('=', Token{COMP, NOTEQUAL}, Token{LOGOP, 0})
	case cur == '>':
		l.handleDoubleSingle('=', Token{COMP, GEQ}, Token{COMP, GREATER})
	case cur == '<':
		l.handleDoubleSingle('=', Token{COMP, LEQ}, Token{COMP, LESS})
	case cur == '&':
		l.handleDoubleSingle('&', Token{LOGOP, AND}, Token{LOGOP, AND})
	case cur == '|':
		l.handleDoubleSingle('|', Token{LOGOP, OR}, Token{LOGOP, OR})
	case IsNumeric(cur):
		err = l.handleNumeric()
	case IsString(cur):
		l.handleString()
	case IsSpace(cur):
	default:
		return errors.New("unrecognized symbol: " + string(cur))
	}
	return err
}

// add adds the token to the result
func (l *lexer) add(t Token) {
	l.res = append(l.res, t)
}

// untilFalse returns the runes until f becomes false
func (l *lexer) untilFalse(f func(rune) bool) []rune {
	ret := []rune{}
	var i int
	for i = l.pos; i < l.len; i++ {
		r := l.data[i]
		if !f(r) {
			break
		}
		ret = append(ret, r)
	}
	l.pos = i - 1
	return ret
}

// handleNumeric handles a situation where the next token is numeric
func (l *lexer) handleNumeric() error {
	str := l.untilFalse(IsNumeric)
	f, err := strconv.ParseFloat(string(str), 64)
	if err != nil {
		return err
	}
	l.add(Token{NUM, f})
	return nil
}

// handleString handles a situation where the next token is a func/var/bool
func (l *lexer) handleString() {
	str := l.untilFalse(IsString)
	l.res = append(l.res, stringToken(str, l.ctx)...)
}

// handleSingleOrDouble handles a situation like the > symbol where the token could either be > or >=.
func (l *lexer) handleDoubleSingle(next rune, double, single Token) {
	if l.pos+1 < l.len {
		if l.data[l.pos+1] == next {
			l.add(double)
			l.pos++
			return
		}
	}
	l.add(single)
}

// stringToken makes tokens from the given string and context
func stringToken(r []rune, ctx *Context) []Token {
	s := string(r)
	if strings.ToLower(s) == "true" {
		return []Token{{BOOL, true}}
	}
	if strings.ToLower(s) == "false" {
		return []Token{{BOOL, false}}
	}
	for n := range ctx.Funcs {
		if n == s {
			return []Token{{FUNC, n}}
		}
	}
	for n := range ctx.Vars {
		if n == s {
			return []Token{{VAR, n}}
		}
	}
	return []Token{{VAR, s}}
}
