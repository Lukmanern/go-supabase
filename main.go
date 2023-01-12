package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"go-supabase/banner"
	"go-supabase/database"
	"go-supabase/handler"
	"go-supabase/helper"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Fields for storing TODO
type Todos struct {
	Id 				 uint64
	Todo, Status, Created_at string
	Deleted_at 			 sql.NullString
}

var db	= database.DatabaseConnection()
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var userInput, index, status uint64
	var todo string
	var err error
	var todoStatus = []string{"done", "inprogress", "todo"}

	banner.ShowBanner()
	for {
		showOptions()
		userInput, err = strconv.ParseUint(helper.GetUserInput("Option : "), 10, 64)
		handler.CheckError(err)

		switch userInput{
		case 0:
			fmt.Println("\n\nSee u ^-^")
			os.Exit(1)
		case 1:
			fmt.Println("> Showing Todos (without Trahsed)")
			get("SELECT * FROM todos WHERE deleted_at IS NULL ORDER BY id")
		case 2:
			fmt.Println("> Showing All Todos (W/ Trahsed)")
			get("SELECT * FROM todos ORDER BY id")
		case 3:
			fmt.Println("> Showing All Trashed Todos")
			get("SELECT * FROM todos WHERE deleted_at IS NOT NULL ORDER BY id")
		case 4:
			fmt.Println("> Create New Todo")
			todo = helper.GetUserInput("New Todo : ")
			create(todo)
		case 5:
			fmt.Println("> Edit Todo")
			index, err = strconv.ParseUint(helper.GetUserInput("Todo Index : "), 10, 64)
			handler.CheckError(err)
			todo = helper.GetUserInput("New Todo : ")
			update(index, todo)

		case 6:
			fmt.Println("> Update Todo's Status")
			index, err = strconv.ParseUint(helper.GetUserInput("Todo Index : "), 10, 64)
			handler.CheckError(err)
			for i, status := range todoStatus{
				fmt.Printf("%v) %s\n", i, status)
			}
			status, err = strconv.ParseUint(helper.GetUserInput("Status : "), 10, 64)
			handler.CheckError(err)
			if status > 2 {
				fmt.Println("> Update Status Failed, please input 0-2")
				continue
			}
			updateStatus(index, todoStatus[status])

		case 7:
			fmt.Println("> SoftDelete Todo")
			index, err = strconv.ParseUint(helper.GetUserInput("Todo Index : "), 10, 64)
			handler.CheckError(err)
			softDelete(index)

		case 8:
			fmt.Println("> Restore Todo (From SoftDelete)")
			index, err = strconv.ParseUint(helper.GetUserInput("Todo Index : "), 10, 64)
			handler.CheckError(err)
			restore(index)

		case 9:
			fmt.Println("> Destroy (Delete Permanent)")
			index, err = strconv.ParseUint(helper.GetUserInput("Todo Index : "), 10, 64)
			handler.CheckError(err)
			if !helper.VerifyUserAction() {
				fmt.Println("> Destroy Failed, the verify pin is wrong")
				continue
			}
			destroy(index)

		case 10:
			fmt.Println("> Hard Reset (Drop todo-table -> re-create table)")
			if !helper.VerifyUserAction() {
				fmt.Println("> Destroy Failed, the verify pin is wrong")
				continue
			}
			hardReset()

		default:
			fmt.Println("> Please re-input 0 to 10")
		}
	}
}

// prints a list of options to the console.
func showOptions() {
	// Print a header message
	fmt.Println("\nOptions :")
	fmt.Println("0. Exit App					6. Update Todo's Status")
	fmt.Println("1. Show Todos					7. SoftDelete Todo")
	fmt.Println("2. Show All Todos (W/ Trashed Todos)		8. Restore Todo (From SoftDelete)")
	fmt.Println("3. Show Just Trashed Todos			9. Destroy (Delete Permanent)")
	fmt.Println("4. Create New Todo				10. Hard Reset (Drop -> re-create Table)")
	fmt.Println("5. Edit Todo")
}

// get retrieves rows from the "todos" table based on the provided query and prints them to the console.
func get(query string) {
	// Declare variables to store the retrieved 
	// todo data and any errors that may occur
	var todo Todos
	var rows *sql.Rows
	// row representation of the todo data
	var row string
	var err error

	// Execute the query and store the result rows
	rows, err = db.Query(query)

	// Check for any errors that occurred during the query execution
	handler.CheckError(err)

	// Print a header message
	fmt.Println("\nid todo (s:status) (d:deleted_at)")

	// Iterate over the result rows
	for rows.Next() {
		// Scan the current row into the todo variable
		err = rows.Scan(&todo.Id, &todo.Todo, &todo.Status, &todo.Created_at, &todo.Deleted_at)

		// Check for any errors that occurred during the scan
		handler.CheckError(err)

		// row representation of the todo data
		row = fmt.Sprintf("%v) %s (s:%s) ", todo.Id, todo.Todo, todo.Status)

		// If the deleted_at column is not null, include it in the string representation
		if todo.Deleted_at.Valid {
			row += fmt.Sprintf("(d:%s)", todo.Deleted_at.String)
		}

		// Print the string representation of the todo data
		fmt.Println(row)
	}
}

// The checkingRowsAffected function checks if a certain number 
// of rows were affected by a database operation.
// If at least one row was affected, it prints a success message. 
// Otherwise, it prints a failure message.
func checkingRowsAffected(rowsAffected int64, functionName string) {
	// If at least one row was affected...
	if rowsAffected > 0 {
		// ...print a success message.
		fmt.Printf("> Success %s, affect %v rows in database\n", functionName, rowsAffected)
		// stop the function
		return
	}
	
	// If no rows were affected...
	// ...print a failure message.
	fmt.Printf("> Failed to %s, affect 0 row in database", functionName)
}

// The create function inserts a new todo into the database.
func create(todo string) {
	// Initialize variables for storing the SQL 
	// statement, the result of executing the statement,
	// the number of rows affected, and any error that might occur.
	var storeSQL string
	var result sql.Result
	var rowsAffect int64
	var err error
	
	// Construct the SQL statement for inserting the new todo.
	storeSQL 	= fmt.Sprintf("INSERT INTO todos(todo) VALUES ('%s')", todo)
	
	// Execute the SQL statement.
	result, err = db.Exec(storeSQL)
	
	// Check for any errors that might have occurred.
	handler.CheckError(err)
	
	// Get the number of rows affected by the SQL statement.
	rowsAffect, err = result.RowsAffected()
	handler.CheckError(err)
	
	// Check if any rows were affected by the SQL statement.
	checkingRowsAffected(rowsAffect, "Create New Todo")
}

func update(index uint64, newTodo string) {
	// Create the UPDATE SQL statement with the new todo and the index
	updateSQL := fmt.Sprintf("UPDATE todos SET todo = '%s' WHERE id = %v", newTodo, index)

	// Execute the UPDATE statement
	result, err := db.Exec(updateSQL)
	handler.CheckError(err)

	// Get the number of rows affected by the UPDATE statement
	rowsAffect, err := result.RowsAffected()
	handler.CheckError(err)

	// Check the number of rows affected and print a message
	checkingRowsAffected(rowsAffect, "Update Todo")
}

// updateStatus updates the status of a todo item in the "todos" table of the database
func updateStatus(index uint64, status string) {
	// Construct the UPDATE SQL query using the provided status and index values
	updateSQl := fmt.Sprintf("UPDATE todos SET status = '%s' WHERE id = %v", status, index)

	// Execute the update query and store the result
	result, err := db.Exec(updateSQl)
	handler.CheckError(err)

	// Get the number of rows affected by the update query
	rowsAffect, err := result.RowsAffected()
	handler.CheckError(err)

	// Pass the number of affected rows and 
	// a message to the checkingRowsAffected function
	checkingRowsAffected(rowsAffect, "Update Todo's Status")
}

func now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func softDelete(index uint64) {
	// Create the UPDATE SQL statement with the current time and the index
	softDeleteSQL := fmt.Sprintf("UPDATE todos SET deleted_at = '%s' WHERE id = %v", now(), index)

	// Execute the UPDATE statement
	result, err := db.Exec(softDeleteSQL)
	handler.CheckError(err)

	// Get the number of rows affected by the UPDATE statement
	rowsAffect, err := result.RowsAffected()
	handler.CheckError(err)

	// Check the number of rows affected and print a message
	checkingRowsAffected(rowsAffect, "Soft Delete Todo")
}


func restore(index uint64) {
	// Create the UPDATE SQL statement with the index
	restoreSQL  := fmt.Sprintf("UPDATE todos SET deleted_at = NULL WHERE id = %v", index)

	// Execute the UPDATE statement
	result, err := db.Exec(restoreSQL)
	handler.CheckError(err)

	// Get the number of rows affected by the UPDATE statement
	rowsAffect, err := result.RowsAffected()
	handler.CheckError(err)

	// Check the number of rows affected and print a message
	checkingRowsAffected(rowsAffect, "Restore Todo")
}

func destroy(index uint64) {
	// Create the DELETE SQL statement with the index
	destroySQL  := fmt.Sprintf("DELETE FROM todos WHERE id = %v", index)

	// Execute the DELETE statement
	result, err := db.Exec(destroySQL)
	handler.CheckError(err)

	// Get the number of rows affected by the DELETE statement
	rowsAffect, err := result.RowsAffected()
	handler.CheckError(err)

	// Check the number of rows affected and print a message
	checkingRowsAffected(rowsAffect, "Restore Todo")
}


func hardReset() {
	// Create the reset SQL statement
	resetSQL := `
	DROP TABLE IF EXISTS todos;
	DROP TYPE IF EXISTS status_options;

	CREATE TYPE status_options AS ENUM('todo', 'inprogress', 'done');

	CREATE TABLE todos (
		id bigint generated always as identity primary key,
		todo text not null,
		status status_options default 'todo',
		created_at timestamp not null default now(),
		deleted_at timestamp default null 
	);`

	// Execute the reset SQL statement
	_, err := db.Exec(resetSQL)
	handler.CheckError(err)

	// Print a success message
	fmt.Println("> Success: Hard Reset")
}