package lang

type MyInt int

func (m *MyInt) Add(b MyInt) {
	var result = int(*m) + int(b)
	*m = MyInt(result)
}
