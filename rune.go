package eval

import (
	"unicode"
)

// isNumeric returns whether the rune is numeric
func isNumeric(r rune) bool {
	return ('0' <= r && r <= '9') || r == '.'
}

// isString returns whether the rune can be a var, func, or bool literal
func isString(r rune) bool {
	return unicode.IsLetter(r)
}

// isLeft returns whether the rune is a left bracket
func isLeft(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

// isRight returns whether the rune is a right bracket
func isRight(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}

// isSpace returns whether the rune is a space
func isSpace(r rune) bool {
	return unicode.IsSpace(r)
}
