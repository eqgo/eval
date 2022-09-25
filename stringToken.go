package eval

import "golang.org/x/exp/slices"

// varFunc is used to get a sorted list of vars and funcs with longest first. Can either be var or func
type varFunc struct {
	name  string
	isVar bool
}

// stringToken makes tokens from the given string and context
func stringToken(r []rune, ctx *Context) []Token {
	s := string(r)
	// handle bool literals
	if s == "true" {
		return []Token{{BOOL, true}}
	}
	if s == "false" {
		return []Token{{BOOL, false}}
	}
	// handle string that is exactly var or func
	for n := range ctx.Vars {
		if n == s {
			return []Token{{VAR, n}}
		}
	}
	for n := range ctx.Funcs {
		if n == s {
			return []Token{{FUNC, n}}
		}
	}
	// get sorted vars and funcs with longest first
	varFuncs := []varFunc{}
	for n := range ctx.Vars {
		varFuncs = append(varFuncs, varFunc{name: n, isVar: true})
	}
	for n := range ctx.Funcs {
		varFuncs = append(varFuncs, varFunc{name: n, isVar: false})
	}
	slices.SortFunc(varFuncs, func(a, b varFunc) bool { return len(a.name) > len(b.name) })
	// handle combinations like xsinx
	slice, ok := stringTokenRecursive(r, ctx, varFuncs)
	if ok {
		return slice
	}
	return []Token{{VAR, s}}
}

func stringTokenRecursive(r []rune, ctx *Context, varFuncs []varFunc) ([]Token, bool) {
	if len(r) == 0 {
		return []Token{}, true
	}
	for _, vf := range varFuncs {
		t := FUNC
		if vf.isVar {
			t = VAR
		}
		for i := 0; i < len(r); i++ {
			s := string(r[:i+1])
			if s == vf.name {
				res, ok := stringTokenRecursive(r[i+1:], ctx, varFuncs)
				if !ok {
					return nil, false
				}
				return append([]Token{{t, vf.name}}, res...), true
			}
		}
	}
	return nil, false
}
