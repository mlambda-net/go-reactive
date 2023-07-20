package pkg

type Bind func(any) Monad

type Monad interface {
	Bind(Bind) Monad
	Catch(func(error))
	UnWrap() (any, error)
}
