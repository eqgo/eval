package eval

import (
	"fmt"
	"strings"
)

// Token is a token in an expression
type Token struct {
	Type  TokenType
	Value any
}

// Tokens returns the tokens for the given expression string
func Tokens(expr string, ctx *Context) ([]Token, error) {
	l := newLexer(append([]rune(expr), '?'), ctx)
	err := l.lex()
	if err != nil {
		return nil, err
	}
	return l.res, nil
}

func (t Token) String() string {
	return fmt.Sprintf("(%v: %v)", t.Type, t.Value)
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
