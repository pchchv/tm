package tm

type Tape struct {
	Symbol []string
	Head   int
}

func NewTape(Symbols []string) *Tape {
	newT := new(Tape)
	newT.Symbol = Symbols
	return newT
}
