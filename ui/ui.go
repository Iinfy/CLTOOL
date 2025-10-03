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
	fmt.Println(CLTOOL_LOGO)
	fmt.Println("\n\nSelect Database(Enter number):\n1.Postgres\n\n0.Exit")
	fmt.Scan(&action)
	switch action {
	case "1":
		fmt.Println("Enter database credentials(databaseLogin/databasePassword/databaseHost/databasePort/databaseName):")
		fmt.Scan(&databaseCredentials)
		err := database.SaveDatabaseCredentials(databaseCredentials, "postgres")
		if err != nil {
			fmt.Println(err)
			return
		}
		connectionStatus := database.ConnectDatabase()
		if connectionStatus {
			fmt.Println("Succesfully connected to PostgreSQL")
			RequestExecutor()
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
		if strings.HasPrefix(sqlCommand, "SELECT ") {
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
