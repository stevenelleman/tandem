package controller

import (
	"sg/services/public-api/dbi"
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
)

func ListSilos() ([]*models.Silo, error) {
	return dbi.ListSilos()
}

func GetSilo(id string) (*models.Silo, error) {
	return dbi.GetSilo(id)
}

func CreateSilo(s *requests.Silo) error {
	return dbi.CreateSilo(s)
}

func UpdateSilo(s *requests.Silo) error {
	return dbi.UpdateSilo(s)
}

func DeleteSilo(id string) error {
	return dbi.DeleteSilo(id)
}
