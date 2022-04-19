package engine

import (
	"game/db"

	"github.com/kataras/iris/v12/websocket"
)

type NewEngine func() Engine

type Engine interface {
	//Get Game Information
	GameInformation() *db.Game
	//Set Game ID
	SetGameID(int)
	//Get Config Information
	Config() ([]byte, error)
	//Set Config
	SetConfig([]byte) error
	//Get Status
	Status() int
	//Get Gamer Number
	Number() int
	//Get Gamer ID list
	Gamers() []int
	//Join Gamer in room
	Join(int, *websocket.Conn) error
}
