package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func ConnectPostgres(databaseType string, databaseLogin string, databasePassword string,
	databaseHost string, databasePort string, databaseName string) (*pgx.Conn, error) {
	connstr := fmt.Sprintf("%v://%v:%v@%v:%v/%v", databaseType, databaseLogin, databasePassword, databaseHost, databasePort, databaseName)
	connection, err := pgx.Connect(context.Background(), connstr)
	if err != nil {
		return nil, err
	}
	return connection, nil

}
