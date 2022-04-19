package rooms

import (
	"errors"
	"fmt"
	"game/db"
	"math/rand"
	"sync"
	"time"

	"github.com/kataras/iris/v12/websocket"
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

func (m *Manager) NewRoom(gameID, id int) (string, error) {
	var key string
	for {
		key = fmt.Sprintf("%06d", rand.Intn(1000000))
		if _, ok := m.roomStore.Load(key); !ok {
			room, err := newRoom(gameID, id, key)
			if err != nil {
				return "", err
			}
			m.roomStore.Store(
				key,
				room,
			)
			go m.tick(room)
			break
		}
	}
	return key, nil
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

func (m *Manager) Join(key string, id int, conn *websocket.Conn) error {
	v, ok := m.roomStore.Load(key)
	if !ok {
		return errors.New("error room id")
	}
	r, _ := v.(*Room)
	return r.Join(id, conn)
}
