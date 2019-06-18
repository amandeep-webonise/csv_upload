package database

import (
	"database/sql"
	"fmt"
)

// PGSQLDBConnectionInitialiser contains postgres database object and configurations
type PGSQLDBConnectionInitialiser struct {
	DB *sql.DB
	*DBConfig
}

// InitialiseConnection initialises postgres connection with specified configurations
func (dw *PGSQLDBConnectionInitialiser) InitialiseConnection() error {
	dw.Driver = `postgres`

	db, err := dw.DBConfig.openConnWithDriver(dw.psqlConnURL())
	if err != nil {
		return err
	}
	dw.DB = db
	return nil
}

// psqlConnURL prepares postgres connection string
func (dw *PGSQLDBConnectionInitialiser) psqlConnURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s", `postgresql`, dw.Username, dw.Password, dw.Host, dw.Port, dw.DatabaseName, dw.ConnectionParams)
}
