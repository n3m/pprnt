package pprnt

var _state *localState = &localState{
	Level: 0,
}

func Print(arg interface{}) string {
	_state.Reset()

	return _ProcessArchitecture(arg, nil, nil)
}
