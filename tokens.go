package eval

import (
	"strconv"
)

// Token is a token in an expression
type Token struct {
	Type  TokenType
	Value any
}

// Tokens returns the tokens for the given expression string
func Tokens(expr string) ([]Token, error) {
	res := []Token{}
	data := append([]rune(expr), '?')
	for pos := 0; pos < len(data); pos++ {
		r := data[pos]
		switch {
		case IsLeft(r):
			res = append(res, Token{LEFT, nil})
		case IsRight(r):
			res = append(res, Token{RIGHT, nil})
		case r == '+':
			res = append(res, Token{ADD, nil})
		case r == '-':
			res = append(res, Token{SUB, nil})
		case r == '*':
			res = append(res, Token{MUL, nil})
		case r == '/':
			res = append(res, Token{DIV, nil})
		case r == '%':
			res = append(res, Token{MOD, nil})
		case IsNumeric(r):
			str := []rune{}
			for i := pos; i < len(data); i++ {

				if !IsNumeric(data[i]) {
					f, err := strconv.ParseFloat(string(str), 64)
					if err != nil {
						return nil, err
					}
					res = append(res, Token{NUM, f})
					pos = i - 1
					break
				}
				str = append(str, data[i])
			}

		}
	}
	return res, nil
}
