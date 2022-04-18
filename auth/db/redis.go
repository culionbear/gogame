package db

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Addr		string			`json:"addr"`
	Password	string			`json:"password,omitempty"`
}

type RedisManager struct {
	handler		*redis.Client
}

func NewRedisManager(config *Config) (*RedisManager, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr: config.Addr,
			Password: config.Password,
		},
	)
	return &RedisManager{
		handler: client,
	}, client.Ping(context.Background()).Err()
}

func (m *RedisManager) Push(key, value string, t time.Duration) error {
	return m.handler.SetEX(context.Background(), key, value, t).Err()
}

func (m *RedisManager) Judge(key, value string) error {
	msg, err := m.handler.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}
	if msg == "" {
		return errors.New(key + " is not found")
	}
	if msg != value {
		return errors.New(value + " is error")
	}
	m.handler.Del(context.Background(), key)
	return nil
}
