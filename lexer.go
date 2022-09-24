package eval

import (
	"errors"
	"strconv"
)

type lexer struct {
	src []rune
	pos int
	len int
	tok []token
	ctx *Context
}

func newLexer(src []rune, ctx *Context) *lexer {
	return &lexer{src: src, pos: 0, len: len(src), tok: []token{}, ctx: ctx}
}

// lex goes through data and sets tok to result
func (l *lexer) lex() error {
	if len(l.src) == 0 {
		return errors.New("lexer: expression string must not be empty")
	}
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
	cur := l.src[l.pos]
	var err error
	switch {
	case isLeft(cur):
		l.add(token{LEFT, nil})
	case isRight(cur):
		l.add(token{RIGHT, nil})
	case cur == '+':
		if l.pos == 0 || isLeft(l.src[l.pos-1]) {
			break
		}
		l.add(token{NUMOP, ADD})
	case cur == '-':
		if l.pos == 0 || isLeft(l.src[l.pos-1]) {
			l.add(token{NUMPRE, NEG})
			break
		}
		l.add(token{NUMOP, SUB})
	case cur == '*':
		l.add(token{NUMOP, MUL})
	case cur == '/':
		l.add(token{NUMOP, DIV})
	case cur == '^':
		l.add(token{NUMOP, POW})
	case cur == '%':
		l.add(token{NUMOP, MOD})
	case cur == '=':
		l.add(token{COMP, EQUAL})
	case cur == ',':
		l.add(token{SEP, nil})
	case cur == '!':
		l.handleDoubleSingle('=', token{COMP, NOTEQUAL}, token{LOGPRE, NOT})
	case cur == '>':
		l.handleDoubleSingle('=', token{COMP, GEQ}, token{COMP, GREATER})
	case cur == '<':
		l.handleDoubleSingle('=', token{COMP, LEQ}, token{COMP, LESS})
	case cur == '&':
		l.handleDoubleSingle('&', token{LOGOP, AND}, token{LOGOP, AND})
	case cur == '|':
		l.handleDoubleSingle('|', token{LOGOP, OR}, token{LOGOP, OR})
	case isNumeric(cur):
		err = l.handleNumeric()
	case isString(cur):
		l.handleString()
	case isSpace(cur):
	default:
		return errors.New("unrecognized symbol: " + string(cur))
	}
	return err
}

// add adds the token to the result
func (l *lexer) add(t token) {
	l.tok = append(l.tok, t)
}

// untilFalse returns the runes until f becomes false
func (l *lexer) untilFalse(f func(rune) bool) []rune {
	ret := []rune{}
	var i int
	for i = l.pos; i < l.len; i++ {
		r := l.src[i]
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
	str := l.untilFalse(isNumeric)
	f, err := strconv.ParseFloat(string(str), 64)
	if err != nil {
		return err
	}
	l.add(token{NUM, f})
	return nil
}

// handleString handles a situation where the next token is a func/var/bool
func (l *lexer) handleString() {
	str := l.untilFalse(isString)
	l.tok = append(l.tok, stringToken(str, l.ctx)...)
}

// handleSingleOrDouble handles a situation like the > symbol where the token could either be > or >=.
func (l *lexer) handleDoubleSingle(next rune, double, single token) {
	if l.pos+1 < l.len {
		if l.src[l.pos+1] == next {
			l.add(double)
			l.pos++
			return
		}
	}
	l.add(single)
}

// stringToken makes tokens from the given string and context
func stringToken(r []rune, ctx *Context) []token {
	s := string(r)
	// handle bool literals
	if s == "true" {
		return []token{{BOOL, true}}
	}
	if s == "false" {
		return []token{{BOOL, false}}
	}
	// handle string that is exactly var or func
	for n := range ctx.Vars {
		if n == s {
			return []token{{VAR, n}}
		}
	}
	for n := range ctx.Funcs {
		if n == s {
			return []token{{FUNC, n}}
		}
	}

	// handle combinations like xsinx
	slice, ok := stringTokenRecursive(r, ctx)
	if ok {
		return slice
	}
	return []token{{VAR, s}}
}

func stringTokenRecursive(r []rune, ctx *Context) ([]token, bool) {
	if len(r) == 0 {
		return []token{}, true
	}
	for v := range ctx.Vars {
		for i := 0; i < len(r); i++ {
			s := string(r[:i+1])
			if s == v {
				res, ok := stringTokenRecursive(r[i+1:], ctx)
				if !ok {
					return nil, false
				}
				return append([]token{{VAR, v}}, res...), true
			}
		}
	}
	for f := range ctx.Funcs {
		for i := 0; i < len(r); i++ {
			s := string(r[:i+1])
			if s == f {
				res, ok := stringTokenRecursive(r[i+1:], ctx)
				if !ok {
					return nil, false
				}
				return append([]token{{FUNC, f}}, res...), true
			}
		}
	}
	return nil, false
}

// fixTokens replaces things like NUM VAR with NUM MUL VAR
func (l *lexer) fixTokens() error {
	if len(l.tok) == 0 {
		return errors.New("lexer: syntax error")
	}
	prev := l.tok[0]
	for i := 1; i < len(l.tok); i++ {
		cur := l.tok[i]
		switch {
		// ex: 9x or 7sin or xy or xsin
		case (prev.typ == NUM || prev.typ == VAR) && (cur.typ == VAR || cur.typ == FUNC):
			l.insert(token{NUMOP, MUL}, i)
			i++
		// )(
		case (prev.typ == RIGHT) && (cur.typ == LEFT):
			l.insert(token{NUMOP, MUL}, i)
			i++
		// ex: 3( or x(
		case (prev.typ == NUM || prev.typ == VAR) && (cur.typ == LEFT):
			l.insert(token{NUMOP, MUL}, i)
			i++
		// ex: )x or )5 or )sin
		case (prev.typ == RIGHT) && (cur.typ == VAR || cur.typ == NUM || cur.typ == FUNC):
			l.insert(token{NUMOP, MUL}, i)
			i++
		// ex: sin3 or sinx
		// TODO: correctly handle things like sin3x
		case (prev.typ == FUNC) && (cur.typ == NUM || cur.typ == VAR):
			l.insert(token{LEFT, nil}, i)
			l.insert(token{RIGHT, nil}, i+2)
			i++
		}
		prev = cur
	}
	return nil
}

// insert inserts the given token at the given index
func (l *lexer) insert(t token, i int) {
	if len(l.tok) <= i {
		l.tok = append(l.tok, t)
		return
	}
	l.tok = append(l.tok[:i+1], l.tok[i:]...)
	l.tok[i] = t
}
