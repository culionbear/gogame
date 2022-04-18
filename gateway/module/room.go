package module

import (
	"game/auth"
	"game/gateway"

	"github.com/kataras/iris/v12"
)

type Room struct {}

func init() {
	gateway.Default.AddParty(new(Game))
}

func (m *Room) Register() (string, func(iris.Party)) {
	return "/room",
	func (p iris.Party) {
		p.Use(auth.Default.Service())
		p.Post("/join")
	}
}
