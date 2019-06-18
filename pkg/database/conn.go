package database

import "database/sql"

// DBConnInitialiser implements database connection methods
type DBConnInitialiser interface {
	InitialiseConnection() (*sql.DB, error)
}

// DBConfig contains database configurations
type DBConfig struct {
	Username         string
	Password         string
	Host             string
	DatabaseName     string
	ConnectionParams string
	Driver           string
	Port             string
}

func (d *DBConfig) openConnWithDriver(url string) (*sql.DB, error) {
	u, err := sql.Open(d.Driver, url)
	if err != nil {
		return nil, err
	}
	if err := u.Ping(); err != nil {
		return nil, err
	}
	return u, nil
}
