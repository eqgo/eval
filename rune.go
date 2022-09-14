package eval

// IsNumeric returns whether the rune is numeric
func IsNumeric(r rune) bool {
	return ('0' <= r && r <= '9') || r == '.'

}

// IsLeft returns whether the rune is a left bracket
func IsLeft(r rune) bool {
	return r == '(' || r == '[' || r == '{'
}

// IsRight returns whether the rune is a right bracket
func IsRight(r rune) bool {
	return r == ')' || r == ']' || r == '}'
}
