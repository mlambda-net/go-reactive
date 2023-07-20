package pkg

type just struct {
	value any
}

func (j *just) Bind(bind Bind) Monad {
	return bind(j.value)
}

func (j *just) Catch(handle func(err error)) {
	// this should be empty not error
}

func (j *just) UnWrap() (any, error) {
	return j.value, nil
}
