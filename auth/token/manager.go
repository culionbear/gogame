package token

import (
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v4"
	jardiniere "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Manager struct {
	key     string
	method  *jwt.SigningMethodHMAC
	handler *jardiniere.Middleware
}

func New(config *Config) *Manager {
	m := &Manager{
		key:    config.Key,
		method: jwt.SigningMethodHS256,
	}
	m.handler = jardiniere.New(jardiniere.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(m.key), nil
		},
		SigningMethod: m.method,
		ErrorHandler: func(ctx iris.Context, err error) {
			ctx.StatusCode(iris.StatusMethodNotAllowed)
			ctx.JSON(iris.Map{
				"msg": "用户信息已过期",
				"err": err.Error(),
			})
		},
	})
	return m
}

func (m *Manager) GetKey(id int) (string, error) {
	table := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Duration(7*24) * time.Hour).Unix(),
		"iss":  "DG",
		"id":   id,
		"rand": rand.Int(),
	}
	return jwt.NewWithClaims(m.method, table).SignedString([]byte(m.key))
}

func (m *Manager) Service() func(iris.Context) {
	return m.handler.Serve
}

func (m *Manager) GetValue(ctx iris.Context, key string) any {
	return ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)[key]
}
