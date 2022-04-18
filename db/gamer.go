package db

import "game/auth"

func (m *Manager) AddGamer(name, phone string) error {
	_, err := m.handler.Exec(
		"insert into gamer (name,phone) values (?,?)",
		name,
		phone,
	)
	return err
}

func (m *Manager) GetGamerIDWithPhone(phone string) (int, error) {
	var id int
	err := m.handler.QueryRow(
		"select id from gamer where phone = ?",
		phone,
	).Scan(&id)
	return id, err
}

func (m *Manager) GetGamerIDWithPassword(user, password string) (int, error) {
	password = auth.Default.MD5(password)

	var id int
	err := m.handler.QueryRow(
		"select id from gamer where (name = ? or phone = ?) and password = ?",
		user,
		user,
		password,
	).Scan(&id)
	return id, err
}

func (m *Manager) UpdatePassword(id int, password string) error {
	password = auth.Default.MD5(password)

	_, err := m.handler.Exec(
		"update gamer set password = ? where id = ?",
		password,
		id,
	)
	return err
}

func (m *Manager) UpdateName(id int, name string) error {
	_, err := m.handler.Exec(
		"update gamer set name = ? where id = ?",
		name,
		id,
	)
	return err
}

func (m *Manager) ExistsGamerWithPhone(phone string) bool {
	var count int
	m.handler.QueryRow(
		"select count(id) from gamer where phone = ?",
		phone,
	).Scan(&count)
	return count != 0
}

func (m *Manager) ExistsGamerWithName(name string) bool {
	var count int
	m.handler.QueryRow(
		"select count(id) from gamer where name = ?",
		name,
	).Scan(&count)
	return count != 0
}

func (m *Manager) GetGamerName(id int) (string, error) {
	var name string
	err := m.handler.QueryRow(
		"select name from gamer where id = ?",
		id,
	).Scan(&name)
	return name, err
}

func (m *Manager) GetGamerPhone(id int) (string, error) {
	var phone string
	err := m.handler.QueryRow(
		"select phone from gamer where id = ?",
		id,
	).Scan(&phone)
	return phone, err
}
