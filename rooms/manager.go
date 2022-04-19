package rooms

import (
	"errors"
	"fmt"
	"game/db"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Manager struct {
	roomStore	*sync.Map
	gamerStore	*sync.Map
}

var Default = New()

func New() *Manager {
	return &Manager{
		roomStore: new(sync.Map),
		gamerStore: new(sync.Map),
	}
}

func (m *Manager) NewRoom(infor db.Game, id int) string {
	var key string
	for {
		key = fmt.Sprintf("%06d", rand.Intn(1000000))
		if _, ok := m.roomStore.Load(key); !ok {
			room := newRoom(infor, id, key)
			m.roomStore.Store(
				key,
				room,
			)
			go m.tick(room)
			break
		}
	}
	return key
}

func (m *Manager) GetRoomInformation(key string) (db.Game, int, error) {
	v, ok := m.roomStore.Load(key)
	if !ok {
		return db.Game{}, 0, errors.New("error room id")
	}
	r, _ := v.(*Room)
	return r.information, r.GetNumber(), nil
}

func (m *Manager) tick(room *Room) {
	ch := room.signal
	for {
		key := <- ch
		gamers := room.GetGamers()
		for _, id := range gamers {
			if v, ok := m.gamerStore.Load(id); ok {
				r, _ := v.(*Room)
				if r.id == key {
					m.gamerStore.Delete(id)
				}
			}
		}
		m.roomStore.Delete(key)
		break
	}
}

func (m *Manager) Judge(key string, id int) bool {
	v, ok := m.roomStore.Load(key)
	if !ok {
		return false
	}
	r, _ := v.(*Room)
	return r.Judge(id)
}
