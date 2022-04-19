package rooms

import (
	"game/db"
	"game/engine"
	"game/engine/options"
	"time"

	"github.com/kataras/iris/v12/websocket"
)

type Room struct {
	//game information [const]
	information	db.Game
	//game engine
	engine		options.Engine
	//room admin id [const]
	admin		int
	//room id in rooms store [const]
	id			string
	//send close signal to rooms manager
	signal		chan string
	//send close when game over
	closer		chan bool
}

func newRoom(gameID, admin int, id string) (*Room, error) {
	e, err := engine.Default.NewGameEngine(gameID)
	if err != nil {
		return nil, err
	}
	m := &Room{
		information: *e.GameInformation(),
		engine: e,
		admin: admin,
		id: id,
		signal: make(chan string),
		closer: make(chan bool),
	}
	go m.close()
	return m, nil
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
	return m.engine.Gamers()
}

func (m *Room) GetNumber() int {
	return m.engine.Number()
}

func (m *Room) GetStatus() int {
	return m.engine.Status()
}

func (m *Room) Join(id int, conn *websocket.Conn) error {
	return m.engine.Join(id, conn)
}

func (m *Room) SetConfig(id int, buf []byte) (bool, error) {
	if id != m.admin {
		return false, nil
	}
	return true, m.engine.SetConfig(buf)
}

func (m *Room) GetConfig() ([]byte, error) {
	return m.engine.Config()
}
