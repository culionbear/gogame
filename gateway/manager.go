package gateway

import (
	"game/rooms"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/rs/cors"
)

type Manager struct {
	RoomHandler    *rooms.Manager
	GatewayHandler *iris.Application
}

var Default = New()

func New() *Manager {
	return &Manager{
		RoomHandler:    rooms.New(),
		GatewayHandler: iris.New(),
	}
}

func (m *Manager) Run(addr string) error {
	m.loadCors()
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

func (m *Manager) loadCors() {
	c := cors.New(
		cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{http.MethodPost, http.MethodGet, http.MethodPut},
			MaxAge:           3600,
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		},
	)
	m.GatewayHandler.WrapRouter(c.ServeHTTP)
}
