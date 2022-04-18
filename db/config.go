package db

import "fmt"

type Config struct {
	Name		string	`json:"name"`
	Password	string	`json:"password"`
	Url			string	`json:"url"`
	DB			string	`json:"db"`
}

func (m *Config) String() string {
	return 	fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		m.Name,
		m.Password,
		m.Url,
		m.DB,
	)
}
