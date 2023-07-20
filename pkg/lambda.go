package pkg

type LambdaObserver interface {
	Observer
	WithError(func(err error)) LambdaObserver
	WithSuccess(func(value any)) LambdaObserver
	WithCompleted(func()) LambdaObserver
}

type lambdaObserver struct {
	next     func(value any)
	complete func()
	errors   func(err error)
}

func (o *lambdaObserver) OnNext(value any) {
	o.next(value)
}

func (o *lambdaObserver) OnComplete() {
	o.complete()
}

func (o *lambdaObserver) OnError(err error) {
	o.errors(err)
}

func (o *lambdaObserver) WithSuccess(onSuccess func(value any)) LambdaObserver {
	o.next = onSuccess
	return o
}

func (o *lambdaObserver) WithCompleted(onComplete func()) LambdaObserver {
	o.complete = onComplete
	return o
}

func (o *lambdaObserver) WithError(onError func(err error)) LambdaObserver {
	o.errors = onError
	return o
}

func OnSuccess(success func(value any)) LambdaObserver {
	return &lambdaObserver{
		next: success,
		complete: func() {
			// success
		},
		errors: func(_ error) {
			// error
		},
	}
}
