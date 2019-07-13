package lang

import (
	"fmt"
	"testing"
)

func TestMyInt(t *testing.T) {
	var a MyInt = 1
	var b MyInt = 2
	a.Add(b)
	fmt.Println(a, b)
}
