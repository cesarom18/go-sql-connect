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

// Create Client
func CreateClient(client models.Client) {
	connection.ConnectDB()
	defer connection.CloseDB()

	query := "INSERT INTO clients VALUES (null, ?, ?, ?)"
	_, err := connection.Db.Exec(query, client.Name, client.Email, client.Phone)
	// Check Query Errors
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Registered Client")
}

// Update Client
func UpdateClient(id int, client models.Client) {
	connection.ConnectDB()
	defer connection.CloseDB()

	query := "UPDATE clients SET name=?, email=?, phone=? WHERE id=?"
	_, err := connection.Db.Exec(query, client.Name, client.Email, client.Phone, id)
	// Check Query Errors
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Updated Client")
}

// Delete Client
func DeleteClient(id int) {
	connection.ConnectDB()
	defer connection.CloseDB()

	query := "DELETE FROM clients WHERE id=?"
	_, err := connection.Db.Exec(query, id)
	// Check Query Errors
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Deleted Client")
}
