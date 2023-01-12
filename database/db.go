package database

import (
	"database/sql"
	"errors"
	"fmt"
	"go-supabase/handler"
	"os"

	"github.com/joho/godotenv"
)

// getENV retrieves environment variables from a .env file
// at the specified path and returns them in a map.
func getENV() map[string]string {
	// Create an empty map to store 
	// the environment variables
	result := make(map[string]string)

	// A slice of keys for the environment 
	// variables we want to retrieve
	keys := []string{"dbname", "user", "port", "host", "password"}

	// The path to the .env file
	path := "C:/Users/Lenovo/go/src/DB_CLI/database/.env"

	// Check if the .env file exists
	_, err := os.Stat(path)
	handler.CheckError(err)

	// Load the environment variables 
	// from the .env file
	godotenv.Load(path)

	// Iterate over the keys and retrieve
	// the corresponding environment variables
	for _, key := range keys {
		result[key], _ = os.LookupEnv(key)
		// If the environment variable is an empty string,
		// there was an error retrieving it
		if result[key] == "" {
			// Modify the key to include 
			// a message about the error
			key = "typo in key="+key
			handler.CheckError(errors.New(key))
		}
	}

	// Return the map of environment variables
	return result
}

// connects to a PostgreSQL database using the specified 
// connection parameters and returns the connection.
func DatabaseConnection() *sql.DB {
	// Declare variables to store the connection 
	// and any errors that may occur
	var connection *sql.DB
	var err error

	// Get the connection parameters from the .env file
	env := getENV()

	// Construct the connection string using the retrieved parameters
	param := fmt.Sprintf("connect_timeout=20 user=%s host=%s port=%s dbname=%s password=%s",
		env["user"], env["host"], env["port"], env["dbname"], env["password"])

	// Open a connection to the database 
	// using the connection string
	connection, err = sql.Open("postgres", param)
	handler.CheckError(err)

	// Return the connection
	return connection
}