package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/amirthapa27/precize-test/database"
	"github.com/amirthapa27/precize-test/handler"
)

func main() {

	database.ConnectDB()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		if !scanner.Scan() {
			break
		}
		option := strings.TrimSpace(scanner.Text())
		switch option {
		case "1":
			handler.InsertData(scanner)
		case "2":
			handler.ViewAllData()
		case "3":
			handler.GetRank(scanner)
		case "4":
			handler.UpdateScore(scanner)
		case "5":
			handler.DeleteOne(scanner)
		case "q":
			return
		default:
			fmt.Println("")
			fmt.Println("***Invalid option. Please try again.***")
			fmt.Println("")
		}
	}
}

func printMenu() {
	fmt.Println("MENU:")
	fmt.Println("1. Insert data")
	fmt.Println("2. View all data")
	fmt.Println("3. Get rank")
	fmt.Println("4. Update score")
	fmt.Println("5. Delete one record")
	fmt.Println("q. Quit")
	fmt.Print("Select an option: ")
}
