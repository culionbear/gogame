package db

type Game struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Logo		string	`json:"logo"`
	MaxGamer	int		`json:"max_gamer"`
	MinGamer	int		`json:"min_gamer"`
	Infor		string	`json:"infor"`
	Rule		string	`json:"rule"`
}

func (m *Manager) GetGameList() ([]Game, error) {
	rows, err := m.handler.Query(
		"select id,name,logo,max_gamer,min_gamer,infor,rule from game",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	list := make([]Game, 0)
	for rows.Next() {
		var msg Game
		err = rows.Scan(
			&msg.ID,
			&msg.Name,
			&msg.Logo,
			&msg.MaxGamer,
			&msg.MinGamer,
			&msg.Infor,
			&msg.Rule,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, msg)
	}
	return list, nil
}

func (m *Manager) GetGameInformation(id int) (Game, error) {
	var msg Game
	err := m.handler.QueryRow(
		"select id,name,logo,max_gamer,min_gamer,infor,rule from game where id = ?",
		id,
	).Scan(
		&msg.ID,
		&msg.Name,
		&msg.Logo,
		&msg.MaxGamer,
		&msg.MinGamer,
		&msg.Infor,
		&msg.Rule,
	)
	return msg, err
}

func (m *Manager) CleanGameTable() error {
	_, err := m.handler.Exec(
		"truncate table game",
	)
	return err
}

func (m *Manager) AddGame(infor *Game) error {
	ids, err := m.handler.Exec(
		"insert into game (name,phone) values (?,?)",
		infor.Name,
		infor.Rule,
		infor.Logo,
		infor.Infor,
		infor.MaxGamer,
		infor.MinGamer,
	)
	if err != nil {
		return err
	}
	id, err := ids.LastInsertId()
	infor.ID = int(id)
	return err
}
