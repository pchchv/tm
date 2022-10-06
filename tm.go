package tm

type TM struct {
	Input        *Tape
	States       map[string]bool
	FinalStates  map[string]bool
	Inputs       map[string]bool
	Configs      map[ConfigIn]ConfigOut
	CurrentState string
}

type ConfigIn struct {
	SrcState string
	Input    string
}

type ConfigOut struct {
	DstState    string
	ModifiedVal string
	TapeMove    int
}

func NewTM() *TM {
	newTM := new(TM)
	newTM.States = make(map[string]bool)
	newTM.FinalStates = make(map[string]bool)
	newTM.Inputs = make(map[string]bool)
	newTM.Configs = make(map[ConfigIn]ConfigOut)
	return newTM
}
