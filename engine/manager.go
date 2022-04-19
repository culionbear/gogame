package engine

import (
	"errors"
	"game/db"
)

type Manager struct {
	store	map[int]NewEngine
}

var Default = New()

func Init() error {
	return Default.Register()
}

func New() *Manager {
	return &Manager{
		store: make(map[int]NewEngine),
	}
}

func (m *Manager) plugins() []NewEngine {
	return []NewEngine {

	}
}

func (m *Manager) Register() error {
	engineStarters := m.plugins()
	err := db.Default.CleanGameTable()
	if err != nil {
		return err
	}
	for _, f := range engineStarters {
		e := f()
		i := e.GameInformation()
		err = db.Default.AddGame(i)
		if err != nil {
			return err
		}
		m.store[i.ID] = f
	}
	return nil
}

func (m *Manager) NewGameEngine(id int) (Engine, error) {
	f, ok := m.store[id]
	if !ok {
		return nil, errors.New("game is not exists")
	}
	e := f()
	e.SetGameID(id)
	return e, nil
}
