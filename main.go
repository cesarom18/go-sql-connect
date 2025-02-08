package main

import (
	"fmt"
	"go-sql-connect/handlers"
	"go-sql-connect/models"
	"strconv"
)

func main() {
	var input string

	for {
		fmt.Println("\n-= Go MySQL Manager =-")
		fmt.Println("1.- Get All Clients")
		fmt.Println("2.- Get Client By ID")
		fmt.Println("3.- Create Client")
		fmt.Println("9.- Exit")

		fmt.Print("Option: ")
		fmt.Scanln(&input)

		if input == "1" {
			handlers.GetClients()
		} else if input == "2" {
			var clientID string

			fmt.Print("Client ID: ")
			fmt.Scanln(&clientID)
			id, err := strconv.Atoi(clientID)

			if err != nil {
				fmt.Println("Error: Invalid client ID. Please enter a valid number.")
				continue
			}

			handlers.GetById(id)
		} else if input == "3" {
			var clientName string
			var clientEmail string
			var clientPhone string

			fmt.Print("Client Name: ")
			fmt.Scanln(&clientName)

			fmt.Print("Client Email: ")
			fmt.Scanln(&clientEmail)

			fmt.Print("Client Phone: ")
			fmt.Scanln(&clientPhone)

			client := models.Client{
				Name:  clientName,
				Email: clientEmail,
				Phone: clientPhone,
			}

			handlers.CreateClient(client)

		} else if input == "9" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
