package lang

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	var i I = 3
	fmt.Println(i)
	i.add(2)
	fmt.Println(i)
}
