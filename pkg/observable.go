package pkg

type Observer interface {
	OnNext(value any)
	OnComplete()
	OnError(err error)
}

type Observable interface {
	Subscribe(observer Observer)
}
