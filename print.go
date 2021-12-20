package pprnt

var _state *localState = &localState{
	Level: 0,
}

// var NoColor bool = !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())

func Print(arg interface{}) string {
	_state.Reset()

	return _ProcessArchitecture(arg, nil, nil)
}
