package controller

import (
	"sg/services/public-api/handlers/requests"
	"sg/services/public-api/models"
)

func (c *Controller) ListSilos() ([]*models.Silo, error) {
	return c.Connection.ListSilos()
}

func (c *Controller) GetSilo(id string) (*models.Silo, error) {
	return c.Connection.GetSilo(id)
}

func (c *Controller) CreateSilo(s *requests.Silo) error {
	return c.Connection.CreateSilo(s)
}

func (c *Controller) UpdateSilo(s *requests.Silo) error {
	return c.Connection.UpdateSilo(s)
}

func (c *Controller) DeleteSilo(id string) error {
	return c.Connection.DeleteSilo(id)
}
