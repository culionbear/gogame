package options

import "sync"

type Mu[T any] struct {
	value	T
	mu		*sync.RWMutex
}

func New[T any](v T) *Mu[T] {
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

type Vector struct {
	list	[]int
	point	int
	mu		*sync.RWMutex
}

func NewVector() *Vector {
	return &Vector{
		list: make([]int, 0),
		mu: new(sync.RWMutex),
	}
}

func (m *Vector) Copy() []int {
	list := make([]int, 0)
	m.mu.RLock()
	for _, v := range m.list {
		list = append(list, v)
	}
	m.mu.RUnlock()
	return list
}

func (m *Vector) Exists(id int) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for _, v := range m.list {
		if id == v {
			return true
		}
	}
	return false
}

func (m *Vector) Add(id int) {
	if m.Exists(id) {
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.list = append(m.list, id)
}

func (m *Vector) Len() int {
	m.mu.RLock()
	length := len(m.list)
	m.mu.RUnlock()
	return length
}

func (m *Vector) Del(id int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.list {
		if id == v {
			m.list = append(m.list[: k], m.list[k + 1: ]...)
			break
		}
	}
}

func (m *Vector) Next() int {
	length := m.Len()
	if length == 0 {
		return 0
	}
	m.mu.Lock()
	num := m.list[m.point % length]
	m.point ++
	m.mu.Unlock()

	return num
}
