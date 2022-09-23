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
	p.stg = p.parseFor(0)
	return nil
}

// parseFor goes through tokens looking for the token given by the given priority
func (p *parser) parseFor(pri priority) *Stage {
	if pri == priorityLen {
		return &Stage{eval: litStage((p.src[0]))}
	}
	tokens := priorityTokensMap[pri]
	for _, t := range tokens {
		for p.pos = 0; p.pos < p.len; p.pos++ {
			cur := p.src[p.pos]
			if cur == t {
				pl := newParser(p.src[:p.pos])
				pl.parse()
				pr := newParser(p.src[p.pos+1:])
				pr.parse()
				return &Stage{
					Left:  pl.stg,
					Right: pr.stg,
					eval:  tokenStageEvalMap[t],
				}
			}
		}
	}
	return p.parseFor(pri + 1)
}
