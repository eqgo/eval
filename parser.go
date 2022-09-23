package eval

type parser struct {
	src []Token
	pos int
	len int
	stg *Stage
}

func newParser(src []Token) *parser {
	return &parser{src: src, pos: 0, len: len(src), stg: nil}
}

// parse goes through tokens and sets stg to result
func (p *parser) parse() error {
	return nil
}
