package gateway

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/neffos"
)

type HttpManager interface {
	Register() (string, func(iris.Party))
}

type WSManager interface {
	Register() (string, *neffos.Server)
}
