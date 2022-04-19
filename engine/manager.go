package engine

import (
	"errors"
	"game/db"
	"game/engine/game/tooth"
	"game/engine/options"
)

type Manager struct {
	store map[int]options.NewEngine
}

var Default = New()

func Init() error {
	return Default.Register()
}

func New() *Manager {
	return &Manager{
		store: make(map[int]options.NewEngine),
	}
}

func (m *Manager) plugins() []options.NewEngine {
	return []options.NewEngine{
		tooth.NewEngine,
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

func (m *Manager) NewGameEngine(id int) (options.Engine, error) {
	f, ok := m.store[id]
	if !ok {
		return nil, errors.New("game is not exists")
	}
	e := f()
	e.SetGameID(id)
	return e, nil
}
