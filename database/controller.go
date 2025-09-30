package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var databaseType, databaseLogin, databasePassword, databaseHost, databasePort, databaseName string
var databaseConnection *pgx.Conn

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
		dconn, err := ConnectPostgres(
			databaseType,
			databaseLogin,
			databasePassword,
			databaseHost,
			databasePort,
			databaseName,
		)
		if err != nil {
			fmt.Printf("Error during connection to database : %v\n", err)
			return false
		}
		databaseConnection = dconn
		return true
	default:
		return false
	}
}

func ExecuteQuery(query string) ([]any, error) {
	rows, err := databaseConnection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var values []any
	for rows.Next() {
		vals, err := rows.Values()
		if err != nil {
			return nil, err
		}
		values = append(values, vals)
	}
	return values, nil
}

func Execute(query string) (pgconn.CommandTag, error) {
	tag, err := databaseConnection.Exec(context.Background(), query)
	if err != nil {
		return tag, err
	}
	return tag, nil
}
