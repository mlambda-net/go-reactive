package pkg

import "fmt"

func Just(value any) Monad {
	return &just{value: value}
}

func Empty() Monad {
	return &empty{}
}

func Fail(err error) Monad {
	return &fail{err: err}
}

func FoG(f Bind, g Bind) Bind {
	return func(value any) Monad {
		return f(value).Bind(func(a any) Monad {
			return g(a)
		})
	}
}

func Pretty(value any) Monad {
	return Just(fmt.Sprintf("%v \n", value))
}

func Print(value any) Monad {
	fmt.Printf("%v", value)
	return Empty()
}
