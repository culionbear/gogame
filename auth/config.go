package auth

import (
	"game/auth/db"
	"game/auth/phone"
	"game/auth/token"
)

type Config struct {
	Redis *db.Config    `json:"redis,omitempty"`
	Phone *phone.Config `json:"phone"`
	Token *token.Config `json:"token"`
}
