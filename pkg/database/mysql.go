package database

import (
	"database/sql"
	"fmt"
)

// MySQLDBConnectionInitialiser ontains mysql database object and configurations
type MySQLDBConnectionInitialiser struct {
	DB *sql.DB
	*DBConfig
}

// InitialiseConnection initialises postgres connection with specified configurations
func (dw *MySQLDBConnectionInitialiser) InitialiseConnection() error {
	dw.Driver = `mysql`

	db, err := dw.DBConfig.openConnWithDriver(dw.mysqlConnURL())
	if err != nil {
		return err
	}
	dw.DB = db
	return nil
}

// mysqlConnURL prepares mysql connection url string
func (dw *MySQLDBConnectionInitialiser) mysqlConnURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", dw.Username, dw.Password, dw.Host, dw.DatabaseName, dw.ConnectionParams)
}
