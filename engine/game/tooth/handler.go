package tooth

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Handler struct {
	tooth [28]bool
	point int
	mu    *sync.RWMutex
}

func NewHandler() *Handler {
	m := &Handler{
		mu: new(sync.RWMutex),
	}
	m.Init()
	return m
}

func (m *Handler) Init() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k := range m.tooth {
		m.tooth[k] = true
	}
	m.point = rand.Intn(28)
}

func (m *Handler) Information() [28]bool {
	var list [28]bool
	m.mu.RLock()
	for k, v := range m.tooth {
		list[k] = v
	}
	m.mu.RUnlock()
	return list
}

func (m *Handler) Option(num int) (bool, error) {
	if num >= 28 {
		return false, errors.New("tooth number error")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.tooth[num] {
		return false, errors.New("tooth is used")
	}
	m.tooth[num] = true
	if m.point == num {
		return true, nil
	}
	return false, nil
}
