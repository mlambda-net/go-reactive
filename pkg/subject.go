package pkg

type Subject struct {
	observers []Observer
}

func (s *Subject) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) OnNext(value any) {
	for _, o := range s.observers {
		o.OnNext(value)
	}
}

func (s *Subject) OnComplete() {
	for _, o := range s.observers {
		o.OnComplete()
	}
}

func (s *Subject) OnError(err error) {
	for _, o := range s.observers {
		o.OnError(err)
	}
}
