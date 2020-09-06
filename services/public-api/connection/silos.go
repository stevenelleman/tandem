package connection

import (
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
)

func (c *Connection) ListSilos() ([]*models.Silo, error) {
	// TODO: use query / prepared statement builder
	rows, err := c.db.Query(`SELECT * FROM silos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	silos := []*models.Silo{}

	// TODO: need better technique for mapping rows to object
	for rows.Next() {
		var id string
		var state string
		err = rows.Scan(&id, &state)
		if err != nil {
			return nil, err
		}
		silos = append(silos, &models.Silo{
			Id:    id,
			State: state,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return silos, nil
}

func (c *Connection) GetSilo(id string) (*models.Silo, error) {
	rows, err := c.db.Query(`SELECT * FROM silos WHERE id=$1;`, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id string
		var state string
		err = rows.Scan(&id, &state)
		if err != nil {
			return nil, err
		}
		return &models.Silo{
			Id:    id,
			State: state,
		}, nil
	}

	return nil, nil
}

func (c *Connection) CreateSilo(s *requests.Silo) error {
	_, err := c.db.Exec(`INSERT INTO silos (id, state) VALUES ($1, $2);`, s.Id, s.State)
	return err
}

func (c *Connection) UpdateSilo(s *requests.Silo) error {
	_, err := c.db.Exec(`UPDATE silos SET state=$1 WHERE id=$2;`, s.State, s.Id)
	return err
}

func (c *Connection) DeleteSilo(id string) error {
	_, err := c.db.Exec(`DELETE FROM silos WHERE id=$1;`, id)
	return err
}
