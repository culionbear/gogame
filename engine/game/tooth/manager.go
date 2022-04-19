package tooth

import (
	"encoding/json"
	"errors"
	"game/db"
	"game/engine/options"
	"sync"

	"github.com/kataras/iris/v12/websocket"
)

func NewEngine() options.Engine {
	return &Manager{
		information: &db.Game{
			Name:     "按牙齿",
			MaxGamer: 12,
			MinGamer: 2,
			Logo:     options.Path("tooth"),
			Infor:    "当心鳄鱼咬到你！",
			Rule:     "每人每回合按一次牙齿，若被咬到则为失败",
		},
		config: options.New(Config{
			Round: 16,
		}),
		status:     options.New(options.STATUS_READY),
		gamers:     options.NewVector(),
		gamerConns: new(sync.Map),
		handler:    NewHandler(),
	}
}

type Manager struct {
	information *db.Game
	config      *options.Mu[Config]
	status      *options.Mu[int]
	gamers      *options.Vector
	gamerConns  *sync.Map

	handler *Handler
}

func (m *Manager) GameInformation() *db.Game {
	return m.information
}

func (m *Manager) SetGameID(id int) {
	m.information.ID = id
}

func (m *Manager) Config() ([]byte, error) {
	return json.Marshal(m.config.Get())
}

func (m *Manager) SetConfig(buf []byte) error {
	if m.status.Get() != options.STATUS_READY {
		return errors.New("must in ready to set config")
	}
	var c Config
	err := json.Unmarshal(buf, &c)
	if err != nil {
		return err
	}
	m.config.Set(c)
	return nil
}

func (m *Manager) Status() int {
	return m.status.Get()
}

func (m *Manager) Number() int {
	return m.gamers.Len()
}

func (m *Manager) Gamers() []int {
	return m.gamers.Copy()
}

func (m *Manager) Join(id int, conn *websocket.Conn) error {
	if m.gamers.Exists(id) {
		if v, ok := m.gamerConns.Load(id); ok {
			c := v.(*websocket.Conn)
			c.Close()
		}
		m.gamerConns.Store(id, conn)
		return nil
	}
	if m.status.Get() != options.STATUS_READY {
		return errors.New("game is starting or ended")
	}
	if m.gamers.Len()+1 > m.information.MaxGamer {
		return errors.New("too many gamers in room")
	}
	m.gamers.Add(id)
	m.gamerConns.Store(id, conn)
	return nil
}

func (m *Manager) Leave(id int) {
	m.gamers.Del(id)
	m.gamerConns.Delete(id)
}

func (m *Manager) Disconnect(id int) {
	m.gamerConns.Delete(id)
}

func (m *Manager) Start() error {
	if m.status.Get() != options.STATUS_READY {
		return errors.New("game is starting or ended")
	}
	if m.gamers.Len() < m.information.MinGamer {
		return errors.New("gamer number is must be greater than min gamer")
	}
	m.status.Set(options.STATUS_PLAYING)

	return nil
}

func (m *Manager) End() error {
	if m.status.Get() != options.STATUS_PLAYING {
		return errors.New("game is not playing")
	}
	m.status.Set(options.STATUS_ENDING)
	return nil
}
