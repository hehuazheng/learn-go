package lang

import (
	"fmt"
	"testing"
)

func TestIntAliasDemo(t *testing.T) {
	var a Month = 3
	var b = a + 2
	outputMonth(a)
	fmt.Println(b)
}
