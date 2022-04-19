package tool

import "github.com/kataras/iris/v12"

func SendMessage(ctx iris.Context, code int, msg string, err ...error) {
	req := iris.Map{
		"msg": msg,
	}
	if len(err) > 0 {
		list := make([]string, 0)
		for _, v := range err {
			if v != nil {
				list = append(list, v.Error())
			}
		}
		req["err"] = list
	}
	ctx.StatusCode(code)
	ctx.JSON(req)
}

func SendBadRequestMessage(ctx iris.Context, err error) {
	SendMessage(ctx, iris.StatusBadRequest, "请求失败", err)
}

func SendBadGatewayMessage(ctx iris.Context, msg string, err error) {
	SendMessage(ctx, iris.StatusBadGateway, msg, err)
}

func SendOKMessage(ctx iris.Context, msg string, infor iris.Map) {
	req := iris.Map{
		"msg": msg,
	}
	if infor != nil {
		req["infor"] = infor
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(req)
}
