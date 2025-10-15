package ui

import (
	"CLTOOL/database"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var redColor = color.New(color.FgRed)
var greenColor = color.New(color.FgGreen)
var cyanColor = color.New(color.FgCyan)
var sprintGreenColor = color.New(color.FgGreen).SprintFunc()

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
		fmt.Println("\n\nSelect Database(Enter number):\n1.Postgres\n2.MySQL\n\n0.Exit")
		fmt.Scan(&action)
		switch action {
		case "1":
			databaseType = "postgres"
		case "2":
			databaseType = "mysql"
		case "0":
			return
		case "exit":
			return
		default:
			redColor.Println("Incorrect database type")
			continue
		}
		if databaseType != "sqlite" {
			redColor.Println("PLEASE DONT USE SPACES")
			greenSlash := sprintGreenColor("/")
			fmt.Printf("%sdatabaseLogin%sdatabasePassword%sdatabaseHost%sdatabasePort%sdatabaseName%s:\n",
				sprintGreenColor("Enter database credentials("), greenSlash, greenSlash, greenSlash, greenSlash, sprintGreenColor(")"))
			fmt.Scan(&databaseCredentials)
			err := database.SaveDatabaseCredentials(databaseCredentials, databaseType)
			if err != nil {
				redColor.Println(err)
				continue
			}
			connectionStatus := database.ConnectDatabase()
			if connectionStatus {
				greenColor.Println("Succesfully connected to ", databaseType)
				RequestExecutor()
			} else {
				redColor.Println("An error occured while connecting to database")
			}
		}

	}

}

func RequestExecutor() {
	scanner := bufio.NewScanner(os.Stdin)
	var sqlCommand string
	for {
		cyanColor.Print("\rEnter command> ")
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
				redColor.Println(err)
				continue
			} else {
				fmt.Println(TableBuilder(columns, values))
			}
		} else {
			result, err := database.Execute(sqlCommand)
			if err != nil {
				redColor.Println(err)
				continue
			}
			rowsAffected, _ := result.RowsAffected()
			greenColor.Println("Rows affected: ", rowsAffected)
		}

	}
}
