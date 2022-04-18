package module

import (
	"errors"
	"game/auth"
	"game/db"
	"game/gateway"
	"game/gateway/tool"

	"github.com/kataras/iris/v12"
)

type Options struct {}

func init() {
	gateway.Default.AddParty(new(Options))
}

func (m *Options) Register() (string, func(iris.Party)) {
	return "/options",
	func (p iris.Party) {
		p.PartyFunc("/login", func(l iris.Party) {
			l.Post("/normal", loginNormal)
			l.Post("/phone", loginPhone)
		})
		p.PartyFunc("/exists", func(e iris.Party) {
			e.Get("/name/{str}", existsName)
			e.Get("/phone/{str}", existsPhone)
		})
		p.Post("/register", register)
		p.Get("/code/{phone}", getCodeWithPhone)
	}
}

func loginNormal(ctx iris.Context) {
	var msg struct {
		User		string	`json:"user"`
		Password	string	`json:"password"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	id, err := db.Default.GetGamerIDWithPassword(msg.User, msg.Password)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "用户名或密码错误", err)
		return
	}
	token, err := auth.Default.GetKey(id)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "token生成失败", err)
		return
	}
	tool.SendOKMessage(ctx, "登录成功", iris.Map{
		"token": token,
	})
}

func loginPhone(ctx iris.Context) {
	var msg struct {
		Phone	string	`json:"phone"`
		Code	string	`json:"code"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	err := auth.Default.VerifyPhoneWithCode(msg.Phone, msg.Code)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "验证码错误或已过期", err)
		return
	}
	id, err := db.Default.GetGamerIDWithPhone(msg.Phone)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "手机号未注册", err)
		return
	}
	token, err := auth.Default.GetKey(id)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "token生成失败", err)
		return
	}
	tool.SendOKMessage(ctx, "登录成功", iris.Map{
		"token": token,
	})
}

func register(ctx iris.Context) {
	var msg struct {
		Name	string	`json:"name"`
		Phone	string	`json:"phone"`
		Code	string	`json:"code"`
	}
	if err := ctx.ReadJSON(&msg); err != nil {
		tool.SendBadRequestMessage(ctx, err)
		return
	}
	if msg.Name == "" {
		tool.SendBadRequestMessage(ctx, errors.New("name is empty"))
		return
	}
	err := auth.Default.VerifyPhoneWithCode(msg.Phone, msg.Code)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "验证码错误或已过期", err)
		return
	}
	err = db.Default.AddGamer(msg.Name, msg.Phone)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "手机号已注册或用户名已存在", err)
		return
	}
	tool.SendOKMessage(ctx, "注册成功", nil)
}

func existsName(ctx iris.Context) {
	str := ctx.Params().Get("str")
	if str == "" {
		tool.SendBadRequestMessage(ctx, errors.New("empty string"))
		return
	}
	tool.SendOKMessage(ctx, "查询成功", iris.Map{
		"flag": db.Default.ExistsGamerWithName(str),
	})
}

func existsPhone(ctx iris.Context) {
	str := ctx.Params().Get("str")
	if str == "" {
		tool.SendBadRequestMessage(ctx, errors.New("empty string"))
		return
	}
	tool.SendOKMessage(ctx, "查询成功", iris.Map{
		"flag": db.Default.ExistsGamerWithPhone(str),
	})
}

func getCodeWithPhone(ctx iris.Context) {
	phone := ctx.Params().Get("phone")
	if phone == "" {
		tool.SendBadRequestMessage(ctx, errors.New("empty phone"))
		return
	}
	err := auth.Default.SendCode(phone)
	if err != nil {
		tool.SendBadGatewayMessage(ctx, "验证码发送失败", err)
		return
	}
	tool.SendOKMessage(ctx, "发送成功", nil)
}
