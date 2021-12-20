package pprnt

type localState struct {
	Level       int
	InOperation bool
}

func (ls *localState) Reset() {
	ls.Level = 0
}

type PrintOptions struct {
	EndComma bool
}
