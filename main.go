package main

import (
	"fmt"
	"go-sql-connect/handlers"
	"strconv"
)

func main() {
	var input string

	for true {
		fmt.Println("\n-= Go MySQL Manager =-")
		fmt.Println("1.- Get All Clients")
		fmt.Println("2.- Get Client By ID")
		fmt.Println("9.- Exit")

		fmt.Print("Option: ")
		fmt.Scanln(&input)

		if input == "1" {
			handlers.GetClients()
		} else if input == "2" {
			fmt.Print("Client ID: ")
			var clientID string
			fmt.Scanln(&clientID)

			id, err := strconv.Atoi(clientID)
			if err != nil {
				fmt.Println("Error: Invalid client ID. Please enter a valid number.")
				continue
			}

			handlers.GetById(id)
		} else if input == "9" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
