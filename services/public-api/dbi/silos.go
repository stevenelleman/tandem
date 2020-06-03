package dbi

import (
	"sg/services/public-api/db"
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
)

func ListSilos() ([]*models.Silo, error) {
	// TODO: use query / prepared statement builder
	rows, err := db.DB.Query(`SELECT * FROM silos`)
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
			Id: id,
			State: state,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return silos, nil
}

func GetSilo(id string) (*models.Silo, error) {
	rows, err := db.DB.Query(`SELECT * FROM silos WHERE id=$1;`, id)
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

func CreateSilo(s *requests.Silo) error {
	_, err := db.DB.Exec(`INSERT INTO silos (id, state) VALUES ($1, $2);`, s.Id, s.State)
	return err
}

func UpdateSilo(s *requests.Silo) error {
	_, err := db.DB.Exec(`UPDATE silos SET state=$1 WHERE id=$2;`, s.State, s.Id)
	return err
}

func DeleteSilo(id string) error {
	_, err := db.DB.Exec(`DELETE FROM silos WHERE id=$1;`, id)
	return err
}