package module

import (
	"encoding/json"
	"errors"
	"game/auth"
	"game/gateway"
	"game/gateway/tool"
	"game/rooms"
	"net/http"

	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gorilla"
)

type WS struct {}

func init() {
	gateway.Default.AddWS(new(WS))
}

func (m *WS) Register() (string, *neffos.Server) {
	ws := websocket.New(
		gorilla.Upgrader(
			gorillaWs.Upgrader{
				CheckOrigin: func(*http.Request) bool {
						return true
				},
			},
		),
		websocket.Events{
			websocket.OnNativeMessage: onChat,
		},
	)
	ws.OnConnect = onConnect
	ws.OnDisconnect = onDisconnect
	return "/ws", ws
}

func onChat(nsConn *websocket.NSConn, msg websocket.Message) error {
	return nil
}

func onConnect(conn *websocket.Conn) error {
	infor, err := readInformation(conn)
	if err != nil {
		tool.WriteWsErrorMessage(conn, err)
		conn.Close()
		return err
	}
	if err = rooms.Default.Join(infor.Room, infor.ID, conn); err != nil {
		tool.WriteWsErrorMessage(conn, err)
		conn.Close()
		return err
	}
	return nil
}

func onDisconnect(conn *websocket.Conn) {
	
}

func readInformation(c *websocket.Conn) (Information, error) {
	var msg Information
	str, err := getToken(c)
	if err != nil {
		return msg, err
	}
	buf, err := auth.Default.AesDecrypt(str)
	if err != nil {
		return msg, err
	}
	err = json.Unmarshal(buf, &msg)
	return msg, err
}

func getToken(c *websocket.Conn) (string, error) {
	token := c.Socket().Request().FormValue("token")
	if token == "" {
		return "", errors.New("token is empty")
	}
	return token, nil
}
