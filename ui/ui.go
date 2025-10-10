package ui

import (
	"CLTOOL/database"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const CLTOOL_LOGO string = `
░█████╗░██╗░░░░░████████╗░█████╗░░█████╗░██╗░░░░░
██╔══██╗██║░░░░░╚══██╔══╝██╔══██╗██╔══██╗██║░░░░░
██║░░╚═╝██║░░░░░░░░██║░░░██║░░██║██║░░██║██║░░░░░
██║░░██╗██║░░░░░░░░██║░░░██║░░██║██║░░██║██║░░░░░
╚█████╔╝███████╗░░░██║░░░╚█████╔╝╚█████╔╝███████╗
░╚════╝░╚══════╝░░░╚═╝░░░░╚════╝░░╚════╝░╚══════╝`

func StartUI() {
	var action string
	var databaseCredentials string
	var databaseType string
	fmt.Println(CLTOOL_LOGO)
	for {
		fmt.Println("\n\nSelect Database(Enter number):\n1.Postgres\n2.MySQL\n3.SQLite\n\n0.Exit")
		fmt.Scan(&action)
		switch action {
		case "1":
			databaseType = "postgres"
		case "2":
			databaseType = "mysql"
		case "3":
			databaseType = "sqlite"
		case "0":
			return
		default:
			fmt.Println("Incorrect database type")
			continue
		}
		if databaseType != "sqlite" {
			fmt.Println("PLEASE DONT USE SPACES!!!")
			fmt.Println("Enter database credentials(databaseLogin/databasePassword/databaseHost/databasePort/databaseName):")
			fmt.Scan(&databaseCredentials)
			err := database.SaveDatabaseCredentials(databaseCredentials, databaseType)
			if err != nil {
				fmt.Println(err)
				return
			}
			connectionStatus := database.ConnectDatabase()
			if connectionStatus {
				fmt.Println("Succesfully connected to ", databaseType)
				RequestExecutor()
			} else {
				fmt.Println("An error occured while connecting to database")
			}
		} else {
			fmt.Println("Coming soon")
		}
	}

}

func RequestExecutor() {
	scanner := bufio.NewScanner(os.Stdin)
	var sqlCommand string
	for {
		fmt.Print("Enter command> ")
		if !scanner.Scan() {
			break
		}
		sqlCommand = scanner.Text()
		if len(sqlCommand) <= 0 {
			continue
		}
		if sqlCommand == "exit" {
			break
		}
		if strings.HasPrefix(sqlCommand, "SELECT ") || strings.HasPrefix(sqlCommand, "SHOW ") {
			values, columns, err := database.ExecuteQuery(sqlCommand)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println(TableBuilder(columns, values))
			}
		} else {
			result, err := database.Execute(sqlCommand)
			if err != nil {
				fmt.Println(err)
				continue
			}
			rowsAffected, _ := result.RowsAffected()
			fmt.Println("Rows affected: ", rowsAffected)
		}

	}
}
