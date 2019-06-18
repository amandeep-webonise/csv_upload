package services

import (
	"database/sql"
	"strconv"

	"github.com/webonise/csv_upload/app/models"
)

// FetchAllUsers return all employee
func (b *Service) FetchAllEmployees() ([]*models.Employee, error) {
	return b.Employee.GetAllEmployees()
}

// FetchAllUsers return all employee
func (b *Service) CheckEmployeesExist(row []string) error {
	emp := &models.Employee{}
	id, _ := strconv.Atoi(row[0])
	//mobile, _ := strconv.ParseInt(row[3], 10, 64)
	emp.ID = id
	emp.Name = sql.NullString{String: row[1], Valid: true}
	emp.Email = sql.NullString{String: row[2], Valid: true}
	//emp.Mobile = sql.NullInt64{Int64: mobile, Valid: true}
	emp.Mobile = sql.NullString{String: row[3], Valid: true}
	emp.Country = sql.NullString{String: row[4], Valid: true}

	return b.Employee.UpsertEmployee(emp)
}
