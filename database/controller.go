package database

import (
	"database/sql"
	"fmt"
	"strings"
)

var databaseType, databaseLogin, databasePassword, databaseHost, databasePort, databaseName string
var databaseConnection *sql.DB

func SaveDatabaseCredentials(credentials string, dbType string) error {
	splitedCredentials := strings.Split(credentials, "/")
	if len(splitedCredentials) != 5 {
		return fmt.Errorf("incorrect credentials")
	}

	databaseLogin = splitedCredentials[0]
	databasePassword = splitedCredentials[1]
	databaseHost = splitedCredentials[2]
	databasePort = splitedCredentials[3]
	databaseName = splitedCredentials[4]
	databaseType = dbType
	return nil
}

func ConnectDatabase() bool {
	switch databaseType {
	case "postgres":
		dbconn, err := ConnectPostgres(
			databaseLogin,
			databasePassword,
			databaseHost,
			databasePort,
			databaseName,
		)
		if err != nil {
			fmt.Println(err)
			return false
		}
		databaseConnection = dbconn
		return true
	case "mysql":
		dbconn, err := ConnectMySql(
			databaseLogin,
			databasePassword,
			databaseHost,
			databasePort,
			databaseName,
		)
		if err != nil {
			fmt.Println(err)
			return false
		}
		databaseConnection = dbconn
		return true
	default:
		return false
	}
}

func ExecuteQuery(query string) ([][]string, []string, error) {
	rows, err := databaseConnection.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	var values [][]string
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		valuesRow := make([]string, len(columns))
		valuesRowAddreses := make([]any, len(columns))
		for i := range len(valuesRow) {
			valuesRowAddreses[i] = &valuesRow[i]
		}
		err := rows.Scan(valuesRowAddreses...)
		if err != nil {
			fmt.Println(err)
		}
		values = append(values, valuesRow)
	}
	return values, columns, nil
}

func Execute(query string) (sql.Result, error) {
	result, err := databaseConnection.Exec(query)
	if err != nil {
		return result, err
	}
	return result, nil
}
