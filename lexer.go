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
	switch {
	case isLeft(cur):
		l.add(Token{LEFT, nil})
	case isRight(cur):
		l.add(Token{RIGHT, nil})
	case cur == '+':
		// if it is the first token it is a positive prefix, which can be ignored
		if len(l.tok) == 0 {
			break
		}
		prev := l.tok[len(l.tok)-1]
		// if the previous token is a num, var, or right bracket, this is a plus symbol; otherwise it is a positive prefix which can be ignored
		if prev.Type == NUM || prev.Type == VAR || prev.Type == RIGHT {
			l.add(Token{NUMOP, ADD})
		}
		// it should not be a bool
		if prev.Type == BOOL {
			return errors.New("lexer: must use float64 values with the + symbol")
		}
	case cur == '-':
		// if it is the first token it is a negative prefix
		if len(l.tok) == 0 {
			l.add(Token{NUMPRE, NEG})
			break
		}
		prev := l.tok[len(l.tok)-1]
		// if the previous token is a num, var, or right bracket, this is a minus symbol; otherwise it is a negative prefix
		if prev.Type == NUM || prev.Type == VAR || prev.Type == RIGHT {
			l.add(Token{NUMOP, SUB})
			break
		}
		// it should not be a bool
		if prev.Type == BOOL {
			return errors.New("lexer: must use float64 values with the - symbol")
		}
		l.add(Token{NUMPRE, NEG})
	case cur == '*':
		l.add(Token{NUMOP, MUL})
	case cur == '/':
		l.add(Token{NUMOP, DIV})
	case cur == '^':
		l.add(Token{NUMOP, POW})
	case cur == '%':
		l.add(Token{NUMOP, MOD})
	case cur == '=':
		l.handleDoubleSingle('=', Token{COMP, EQUAL}, Token{COMP, EQUAL})
	case cur == ',':
		l.add(Token{SEP, nil})
	case cur == '≤':
		l.add(Token{COMP, LEQ})
	case cur == '≥':
		l.add(Token{COMP, GEQ})
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
	case isNumeric(cur):
		err := l.handleNumeric()
		if err != nil {
			return err
		}
	case isString(cur):
		l.handleString()
	case isSpace(cur):
	default:
		return errors.New("unrecognized symbol: " + string(cur))
	}
	return nil
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
	str := l.untilFalse(isNumeric)
	f, err := strconv.ParseFloat(string(str), 64)
	if err != nil {
		return err
	}
	l.add(Token{NUM, f})
	return nil
}

// handleString handles a situation where the next token is a func/var/bool
func (l *lexer) handleString() {
	str := l.untilFalse(isString)
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

// fixTokens replaces things like NUM VAR with NUM MUL VAR
func (l *lexer) fixTokens() error {
	if len(l.tok) == 0 {
		return errors.New("lexer: syntax error")
	}
	prev := l.tok[0]
	for i := 1; i < len(l.tok); i++ {
		cur := l.tok[i]
		switch {
		// ex: 9x or 7sin or xy or xsin or x3
		case (prev.Type == NUM || prev.Type == VAR) && (cur.Type == VAR || cur.Type == FUNC || cur.Type == NUM):
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
	return nil
}

// insert inserts the given token at the given index
func (l *lexer) insert(t Token, i int) {
	if len(l.tok) <= i {
		l.tok = append(l.tok, t)
		return
	}
	l.tok = append(l.tok[:i+1], l.tok[i:]...)
	l.tok[i] = t
}
