package eval

import (
	"fmt"
)

// token is a token in an expression
type token struct {
	typ   tokenType
	value any
}

// tokens returns the tokens for the given expression string with the given context
func tokens(expr string, ctx *Context) ([]token, error) {
	l := newLexer([]rune(expr), ctx)
	err := l.lex()
	if err != nil {
		return nil, err
	}
	err = l.fixTokens()
	return l.tok, err
}

func (t token) String() string {
	return fmt.Sprintf("(%v: %v)", t.typ, t.value)
}
