package lang

type Book struct {
	name   string
	author string
}

func tuple(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
