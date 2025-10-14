package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectPostgres(databaseLogin string, databasePassword string,
	databaseHost string, databasePort string, databaseName string) (*sql.DB, error) {
	connstr := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", databaseLogin, databasePassword, databaseName, databaseHost, databasePort)
	connection, err := sql.Open("postgres", connstr)
	if err != nil {
		return nil, err
	}
	if connection.Ping() == nil {
		return connection, nil
	} else {
		return nil, fmt.Errorf("error during database connection")
	}

}
