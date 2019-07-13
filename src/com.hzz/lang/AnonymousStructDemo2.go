package lang

type APerson struct {
	name string
	addr struct {
		province string
		code     string
	}
}

type AddrNoname struct {
	province string
	code     string
}
