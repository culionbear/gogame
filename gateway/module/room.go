package module

import (
	"errors"
	"game/auth"
	"game/gateway"
	"game/gateway/tool"
	"game/rooms"

	"github.com/kataras/iris/v12"
)

type Information struct {
	ID   int    `json:"id"`
	Room string `json:"room"`
	Rand int64  `json:"rand"`
}

type Room struct{}

func init() {
	gateway.Default.AddParty(new(Game))
}

func (m *Room) Register() (string, func(iris.Party)) {
	return "/room",
		func(p iris.Party) {
			p.Use(auth.Default.Service())
			p.Post("/join", join)
			p.Get("/infor/{id}", getRoomInformation)
			p.Get("/config/{id}", getRoomConfig)
			p.Post("/update/config", updateRoomConfig)
		}
}

func join(ctx iris.Context) {

}

func getRoomInformation(ctx iris.Context) {
	id := ctx.Params().Get("id")
	if id == "" {
		tool.SendBadRequestMessage(ctx, errors.New("error id"))
		return
	}
	infor, num, err := rooms.Default.GetRoomInformation(id)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "房间不存在", err)
		return
	}
	tool.SendOKMessage(ctx, "信息查询成功", iris.Map{
		"infor":  infor,
		"number": num,
	})
}

func getRoomConfig(ctx iris.Context) {

}

func updateRoomConfig(ctx iris.Context) {

}
