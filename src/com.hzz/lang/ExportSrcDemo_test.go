package lang

import (
	"fmt"
	"testing"
)

func TestExportInt(t *testing.T) {
	fmt.Println(Pi)
	t.Log("i'm", pi)
	t.Log("i'm", Pi)
}
