package rooms

import (
	"game/db"
	"sync"
	"time"
)

const (
	STATUS_READY	= iota
	STATUS_PLAYING
	STATUS_ENDING
)

type Room struct {
	information	db.Game
	//TODO:set game engine
	engine		any
	number		int
	status		int
	admin		int
	gamers		[]int
	id			string
	signal		chan string
	closer		chan bool
	mu			*sync.RWMutex
}

func newRoom(infor db.Game, admin int, id string) *Room {
	m := &Room{
		information: infor,
		admin: admin,
		id: id,
		gamers: make([]int, 0),
		signal: make(chan string),
		closer: make(chan bool),
		mu: new(sync.RWMutex),
	}
	go m.close()
	return m
}

func (m *Room) close() {
	ticker := time.NewTicker(time.Hour * time.Duration(12))
	for {
		select {
		case <- ticker.C :
			m.signal <- m.id
			break
		case <- m.closer:
			m.signal <- m.id
			break
		}
	}
}

func (m *Room) GetGamers() []int {
	gamers := make([]int, 0)
	m.mu.RLock()
	for _, v := range m.gamers {
		gamers = append(gamers, v)
	}
	m.mu.RUnlock()
	return gamers
}

func (m *Room) GetNumber() int {
	var number int
	m.mu.RLock()
	number = m.number
	m.mu.RUnlock()
	return number
}
