package eval

import (
	"fmt"
)

// Token is a Token in an expression
type Token struct {
	Type  TokenType
	Value any
}

// Tokens returns the tokens for the given expression string with the given context
func Tokens(expr string, ctx *Context) ([]Token, error) {
	l := newLexer([]rune(expr), ctx)
	err := l.lex()
	if err != nil {
		return nil, err
	}
	err = l.fixTokens()
	return l.tok, err
}

func (t Token) String() string {
	return fmt.Sprintf("(%v: %v)", t.Type, t.Value)
}
