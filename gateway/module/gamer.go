package module

import (
	"game/auth"
	"game/db"
	"game/gateway"
	"game/gateway/tool"

	"github.com/kataras/iris/v12"
)

type Gamer struct {}

func init() {
	gateway.Default.AddParty(new(Gamer))
}

func (m *Gamer) Register() (string, func(iris.Party)) {
	return "/gamer",
	func (p iris.Party) {
		p.Use(auth.Default.Service())
		p.Get("/name", getName)
		p.Get("/phone", getPhone)
		p.PartyFunc("/update", func(u iris.Party) {
			u.Post("/password", updatePassword)
			u.Post("/name", updateName)
		})
		p.Get("/code", getCodeWithID)
	}
}

func getName(ctx iris.Context) {
	id := auth.Default.GetID(ctx)
	name, err := db.Default.GetGamerName(int(id))
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息获取失败", err)
		return
	}
	tool.SendOKMessage(ctx, "信息获取成功", iris.Map{
		"name": name,
	})
}

func getPhone(ctx iris.Context) {
	id := auth.Default.GetID(ctx)
	phone, err := db.Default.GetGamerPhone(int(id))
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息获取失败", err)
		return
	}
	if len(phone) == 11 {
		phone = phone[:3] + "****" + phone[7:]
	}
	tool.SendOKMessage(ctx, "信息获取成功", iris.Map{
		"phone": phone,
	})
}

func updatePassword(ctx iris.Context) {
	var msg struct {
		Code		string	`json:"code"`
		Password	string	`json:"password"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	id := auth.Default.GetID(ctx)
	phone, err := db.Default.GetGamerPhone(int(id))
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息获取失败", err)
		return
	}
	err = auth.Default.VerifyPhoneWithCode(phone, msg.Code)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "验证码错误或不存在", err)
		return
	}
	err = db.Default.UpdatePassword(int(id), msg.Password)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息修改失败", err)
		return
	}
	tool.SendOKMessage(ctx, "信息修改成功", nil)
}

func updateName(ctx iris.Context) {
	var msg struct {
		Name	string	`json:"name"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	id := auth.Default.GetID(ctx)
	err := db.Default.UpdateName(int(id), msg.Name)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息修改失败", err)
		return
	}
	tool.SendOKMessage(ctx, "信息修改成功", nil)
}

func getCodeWithID(ctx iris.Context) {
	id := auth.Default.GetID(ctx)
	phone, err := db.Default.GetGamerPhone(int(id))
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "信息获取失败", err)
		return
	}
	err = auth.Default.SendCode(phone)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "验证码发送失败", err)
		return
	}
	tool.SendOKMessage(ctx, "发送成功", nil)
}
