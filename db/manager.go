package db

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Manager struct {
	handler *sql.DB
}

var Default *Manager

func Init(config *Config) error {
	var err error
	Default, err = New(config)
	return err
}

func New(config *Config) (*Manager, error){
	handler, err := sql.Open(
		"mysql",
		config.String(),
	)
	if err != nil{
		return nil, err
	}
	return &Manager{
		handler: handler,
	}, handler.Ping()
}
