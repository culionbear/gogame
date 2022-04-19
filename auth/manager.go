package auth

import (
	"game/auth/db"
	"game/auth/phone"
	"game/auth/token"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixMilli())
}

type Manager struct {
	phone		*phone.Manager
	db			db.Manager
	token		*token.Manager
	config		*Config
}

var Default *Manager

func Init(config *Config) error {
	var err error
	Default, err = New(config)
	return err
}

func New(config *Config) (*Manager, error) {

	dbClient, err := initDB(config.Redis)
	if err != nil {
		return nil, err
	}

	phoneClient, err := phone.New(config.Phone)
	if err != nil {
		return nil, err
	}


	return &Manager{
		db: dbClient,
		phone: phoneClient,
		token: token.New(config.Token),
		config: config,
	}, nil
}

func initDB(config *db.Config) (db.Manager, error) {
	var client db.Manager = db.NewLocalManager()

	if config != nil {
		var err error
		client, err = db.NewRedisManager(config)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}
