package tooth

import "game/db"

func NewEngine() *Manager {
	return &Manager{}
}

type Manager struct {
	information		db.Game
	status			int
	number			int
	
}
