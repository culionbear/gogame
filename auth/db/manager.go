package db

import "time"

type Manager interface {
	Push(string, string, time.Duration) error
	Judge(string, string) error
}
