package eval

import (
	"fmt"
)

// Token is a token in an expression
type Token struct {
	Type  TokenType
	Value any
}

// Tokens returns the tokens for the given expression string
func Tokens(expr string, ctx *Context) ([]Token, error) {
	l := newLexer([]rune(expr), ctx)
	err := l.lex()
	if err != nil {
		return nil, err
	}
	l.fixTokens()
	return l.tok, nil
}

func (t Token) String() string {
	return fmt.Sprintf("(%v: %v)", t.Type, t.Value)
}
