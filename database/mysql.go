package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySql(databaseLogin string, databasePassword string,
	databaseHost string, databasePort string, databaseName string) (*sql.DB, error) {
	connstr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", databaseLogin, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := sql.Open("mysql", connstr)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
