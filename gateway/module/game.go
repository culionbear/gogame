package module

import (
	"game/auth"
	"game/db"
	"game/gateway"
	"game/gateway/tool"
	"game/rooms"

	"github.com/kataras/iris/v12"
)

type Game struct{}

func init() {
	gateway.Default.AddParty(new(Game))
}

func (m *Game) Register() (string, func(iris.Party)) {
	return "/game",
		func(p iris.Party) {
			p.Use(auth.Default.Service())
			p.Get("/list", getGameList)
			p.Post("/create", createGameRoom)
		}
}

func getGameList(ctx iris.Context) {
	list, err := db.Default.GetGameList()
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息获取失败", err)
		return
	}
	tool.SendOKMessage(ctx, "信息获取成功", iris.Map{
		"list": list,
	})
}

func createGameRoom(ctx iris.Context) {
	var msg struct {
		ID int `json:"id"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	id := auth.Default.GetID(ctx)
	key, err := rooms.Default.NewRoom(msg.ID, int(id))
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "房间创建失败", err)
		return
	}
	tool.SendOKMessage(ctx, "房间创建成功", iris.Map{
		"key": key,
	})
}
