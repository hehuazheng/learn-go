package lang

import (
	"testing"
)

func TestInterfaceDemo1(t *testing.T) {
	factory := ShapeFactory{}
	factory.getShape("c").Draw()
	factory.getShape("r").Draw()
	factory.getShape("s").Draw()
}
