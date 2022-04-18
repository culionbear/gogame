package auth

import (
	"time"
)

func (m *Manager) SendCode(phone string) error {
	code, err := m.phone.SendCode(phone)
	if err != nil {
		return err
	}
	m.db.Push(phone, code, time.Minute * 15)
	return nil
}

func (m *Manager) VerifyPhoneWithCode(phone, code string) error {
	return m.db.Judge(phone, code)
}
