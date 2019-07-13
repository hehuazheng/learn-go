package lang

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	var p = Person{"name", "addr", Career{"carrer2"}}
	fmt.Println(p)
}
