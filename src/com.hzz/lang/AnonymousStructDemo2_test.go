package lang

import (
	"fmt"
	"testing"
)

func TestImportInt(t *testing.T) {
	var a = APerson{"name1", AddrNoname{"hz", "bb"}}
	a.name = "name2"
	fmt.Println(a)
}
