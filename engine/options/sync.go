package options

import "sync"

type Mu[T Value] struct {
	value	T
	mu		*sync.RWMutex
}

func New[T Value](v T) *Mu[T] {
	return &Mu[T]{
		value: v,
		mu: new(sync.RWMutex),
	}
}

func (m *Mu[T]) Set(v T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = v
}

func (m *Mu[T]) Get() T {
	m.mu.RLock()
	v := m.value
	m.mu.RUnlock()
	return v
}
