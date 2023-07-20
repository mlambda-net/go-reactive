package pkg

type Mono struct {
	Subject
}

func (m *Mono) Subscribe(observer Observer) {
	m.observers = append(m.observers, observer)
}

func (m *Mono) OnNext(value any) {

	for _, o := range m.observers {
		o.OnNext(value)
	}
}
