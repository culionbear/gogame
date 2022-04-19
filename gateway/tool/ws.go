package tool

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func WriteWsMessage(conn *websocket.Conn, msg iris.Map) {
	buf, _ := json.Marshal(msg)
	conn.Write(websocket.Message{
		Body: buf,
		IsNative: true,
	})
}

func WriteWsErrorMessage(conn *websocket.Conn, err error) {
	buf, _ := json.Marshal(iris.Map{
		"msg": "与房间的连接断开...",
	})
	conn.Write(websocket.Message{
		Body: buf,
		Err: err,
		IsNative: true,
	})
}
