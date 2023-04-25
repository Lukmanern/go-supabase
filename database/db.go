package database

import (
	"database/sql"
	"fmt"
	"go-supabase/handler"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type databaseConfig struct {
	dbname   string
	user     string
	port     string
	host     string
	password string
}

func getENV() map[string]string {
	// Create an empty map to store
	// the environment variables
	result := make(map[string]string)
	keys := []string{"dbname", "user", "port", "host", "password"}
	path := "C:/Users/Lenovo/go/src/DB_CLI/database/.env"

	// Check if the .env file exists
	if _, err := os.Stat(path); err != nil {
		handler.CheckError(err)
	}
	if err := godotenv.Load(path); err != nil {
		handler.CheckError(err)
	}

	// Iterate over the keys and retrieve
	// the corresponding environment variables
	for _, key := range keys {
		if val := os.Getenv(key); val == "" {
			handler.CheckError(fmt.Errorf("typo in key=%s", key))
		} else {
			result[key] = val
		}
	}

	return result
}

func DatabaseConnection() *sql.DB {
	env := getENV()

	config := databaseConfig{
		dbname:   env["dbname"],
		user:     env["user"],
		port:     env["port"],
		host:     env["host"],
		password: env["password"],
	}

	// fmt.Println("config :", config)

	param := fmt.Sprintf("user=%s host=%s port=%s dbname=%s password=%s",
		config.user, config.host, config.port, config.dbname, config.password)
	conn, err := sql.Open("postgres", param)
	handler.CheckError(err)

	conn.SetMaxOpenConns(5)
	conn.SetMaxIdleConns(3)

	return conn
}
