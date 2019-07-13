package lang

import (
	"fmt"
	"testing"
)

func Test_initB(t *testing.T) {
	var b = B{b: 2, A: A{a: 3}}
	fmt.Println("b value ", b.A.a, b.b)
}
