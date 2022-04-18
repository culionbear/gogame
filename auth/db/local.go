package db

import (
	"errors"
	"sync"
	"time"
)

type information struct {
	value	string
	t		time.Time
}

type LocalManager struct {
	store		*sync.Map
}

func NewLocalManager() *LocalManager {
	m := &LocalManager{
		store: new(sync.Map),
	}
	go m.tick()
	return m
}

func (m *LocalManager) tick() {
	ticker := time.NewTicker(time.Minute)
	for {
		now := <- ticker.C
		m.store.Range(
			func(key, value any) bool {
				infor, _ := value.(*information)
				if now.After(infor.t) {
					m.store.Delete(key)
				}
				return true
			},
		)
	}
}

func (m *LocalManager) Push(key, value string, t time.Duration) error {
	m.store.Store(key, &information{
		value: value,
		t: time.Now().Add(t),
	})
	return nil
}

func (m *LocalManager) Judge(key, value string) error {
	v, ok := m.store.Load(key)
	if !ok {
		return errors.New(key + " is not found")
	}
	msg, _ := v.(*information)
	if msg.value != value {
		return errors.New(value + " is error")
	}
	m.store.Delete(key)
	return nil
}
