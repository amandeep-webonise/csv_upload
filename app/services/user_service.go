package services

import (
	"github.com/webonise/csv_upload/app/models"
)

// FetchAllUsers return all users
func (b *Service) FetchAllUsers() ([]*models.User, error) {
	return b.User.GetAllUsers()
}
