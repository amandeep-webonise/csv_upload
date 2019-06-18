package services

import (
	"github.com/webonise/csv_upload/app/models"
	"github.com/webonise/csv_upload/pkg/logger"
)

// Service contains basic dependencies on which services depends
type Service struct {
	Log      logger.Ilogger
	User     models.UserService
	Employee models.EmployeeService
}

// ServiceProvider provides services to controllers
type ServiceProvider interface {
	FetchAllUsers() ([]*models.User, error)
	CheckEmployeesExist(row []string) error
}
