package eval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	fmt.Println(Tokens("32.7%64.000*8"))
}
