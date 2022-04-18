package auth

import "github.com/kataras/iris/v12"

func (m *Manager) GetKey(id int) (string, error) {
	return m.token.GetKey(id)
}

func (m *Manager) Service() func(iris.Context) {
	return m.token.Service()
}

func (m *Manager) GetValue(ctx iris.Context, key string) any {
	return m.token.GetValue(ctx, key)
}

func (m *Manager) GetID(ctx iris.Context) int32 {
	return int32(m.token.GetValue(ctx, "id").(float64))
}
