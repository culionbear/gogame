package gateway

import (
	"game/rooms"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

type Manager struct {
	RoomHandler		*rooms.Manager
	GatewayHandler	*iris.Application
}

var Default = New()

func New() *Manager {
	return &Manager{
		RoomHandler: rooms.New(),
		GatewayHandler: iris.New(),
	}
}

func (m *Manager) Run(addr string) error {
	return m.GatewayHandler.Run(
		iris.Addr(addr),
	)
}

func (m *Manager) AddParty(api HttpManager) {
	m.GatewayHandler.PartyFunc(api.Register())
}

func (m *Manager) AddWS(api WSManager) {
	path, handler := api.Register()
	m.GatewayHandler.Get(path, websocket.Handler(handler))
}
