package lang

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	var b = Book{name: "book1"}
	var s string = tuple(b.author == "", "nil", "not nil").(string)
	fmt.Println("#", b.author, "#", s)
	fmt.Println(b)
}
