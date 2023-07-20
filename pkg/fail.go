package pkg

type fail struct {
	err error
}

func (f *fail) Bind(_ Bind) Monad {
	return &empty{}
}

func (f *fail) UnWrap() (any, error) {
	return nil, f.err
}

func (f *fail) Catch(handle func(err error)) {
	handle(f.err)
}
