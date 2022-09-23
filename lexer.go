package eval

import (
	"errors"
	"strconv"
)

type lexer struct {
	src []rune
	pos int
	len int
	tok []Token
	ctx *Context
}

func newLexer(src []rune, ctx *Context) *lexer {
	return &lexer{src: src, pos: 0, len: len(src), tok: []Token{}, ctx: ctx}
}

// lex goes through data and sets tok to result
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
	cur := l.src[l.pos]
	var err error
	switch {
	case IsLeft(cur):
		l.add(Token{LEFT, nil})
	case IsRight(cur):
		l.add(Token{RIGHT, nil})
	case cur == '+':
		if l.pos == 0 || IsLeft(l.src[l.pos-1]) {
			break
		}
		l.add(Token{NUMOP, ADD})
	case cur == '-':
		if l.pos == 0 || IsLeft(l.src[l.pos-1]) {
			l.add(Token{NUMPRE, NEG})
			break
		}
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
	case cur == ',':
		l.add(Token{SEP, nil})
	case cur == '!':
		l.handleDoubleSingle('=', Token{COMP, NOTEQUAL}, Token{LOGPRE, NOT})
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
	l.tok = append(l.tok, stringToken(str, l.ctx)...)
}

// handleSingleOrDouble handles a situation like the > symbol where the token could either be > or >=.
func (l *lexer) handleDoubleSingle(next rune, double, single Token) {
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
func stringToken(r []rune, ctx *Context) []Token {
	s := string(r)
	// handle bool literals
	if s == "true" {
		return []Token{{BOOL, true}}
	}
	if s == "false" {
		return []Token{{BOOL, false}}
	}
	// handle string that is exactly var or func
	for n := range ctx.Vars {
		if n == s {
			return []Token{{VAR, n}}
		}
	}
	for n := range ctx.Funcs {
		if n == s {
			return []Token{{FUNC, n}}
		}
	}

	// handle combinations like xsinx
	slice, ok := stringTokenRecursive(r, ctx)
	if ok {
		return slice
	}
	return []Token{{VAR, s}}
}

func stringTokenRecursive(r []rune, ctx *Context) ([]Token, bool) {
	if len(r) == 0 {
		return []Token{}, true
	}
	for v := range ctx.Vars {
		for i := 0; i < len(r); i++ {
			s := string(r[:i+1])
			if s == v {
				res, ok := stringTokenRecursive(r[i+1:], ctx)
				if !ok {
					return nil, false
				}
				return append([]Token{{VAR, v}}, res...), true
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
				return append([]Token{{FUNC, f}}, res...), true
			}
		}
	}
	return nil, false
}

// fixTokens replaces things like NUM VAR with NUM MUL VAR
func (l *lexer) fixTokens() {
	prev := l.tok[0]
	for i := 1; i < len(l.tok); i++ {
		cur := l.tok[i]
		switch {
		// ex: 9x or 7sin or xy or xsin
		case (prev.Type == NUM || prev.Type == VAR) && (cur.Type == VAR || cur.Type == FUNC):
			l.insert(Token{NUMOP, MUL}, i)
			i++
		// )(
		case (prev.Type == RIGHT) && (cur.Type == LEFT):
			l.insert(Token{NUMOP, MUL}, i)
			i++
		// ex: 3( or x(
		case (prev.Type == NUM || prev.Type == VAR) && (cur.Type == LEFT):
			l.insert(Token{NUMOP, MUL}, i)
			i++
		// ex: )x or )5 or )sin
		case (prev.Type == RIGHT) && (cur.Type == VAR || cur.Type == NUM || cur.Type == FUNC):
			l.insert(Token{NUMOP, MUL}, i)
			i++
		// ex: sin3 or sinx
		// TODO: correctly handle things like sin3x
		case (prev.Type == FUNC) && (cur.Type == NUM || cur.Type == VAR):
			l.insert(Token{LEFT, nil}, i)
			l.insert(Token{RIGHT, nil}, i+2)
			i++
		}
		prev = cur
	}
}

// insert inserts the given token at the given index
func (l *lexer) insert(t Token, i int) {
	l.tok = append(l.tok[:i+1], l.tok[i:]...)
	l.tok[i] = t
}
