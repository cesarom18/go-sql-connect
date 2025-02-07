package handlers

import (
	"fmt"
	"go-sql-connect/connection"
	"go-sql-connect/models"
	"log"
)

// Get All Clients
func GetClients() {
	connection.ConnectDB()
	defer connection.CloseDB()

	query := "SELECT id, name, email, phone FROM clients ORDER BY id DESC;"
	data, err := connection.Db.Query(query)
	// Check Query Errors
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nClient List")
	fmt.Println("--------------------")
	for data.Next() {
		client := models.Client{}
		err := data.Scan(&client.Id, &client.Name, &client.Email, &client.Phone)
		// Check Scan Error
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s \n", client.Id, client.Name, client.Email, client.Phone)
	}
}

// Get By ID
func GetById(id int) {
	connection.ConnectDB()
	defer connection.CloseDB()

	query := "SELECT id, name, email, phone FROM clients WHERE id=?;"
	data, err := connection.Db.Query(query, id)
	// Check Query Errors
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nClient By ID")
	fmt.Println("--------------------")
	for data.Next() {
		client := models.Client{}
		err := data.Scan(&client.Id, &client.Name, &client.Email, &client.Phone)
		// Check Scan Error
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Name: %s | Email: %s | Phone: %s \n", client.Id, client.Name, client.Email, client.Phone)
	}
}
