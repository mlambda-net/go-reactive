package pkg

type empty struct {
}

func (e *empty) Bind(_ Bind) Monad {
	return &empty{}
}

func (e *empty) UnWrap() (any, error) {
	return nil, nil
}

func (e *empty) Catch(handle func(err error)) {
	// this should be empty not error
}
