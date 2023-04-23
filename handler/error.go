package handler

import (
	"fmt"
	"go-supabase/vars"
	"os"
)

// checkError checks if there is an error and prints it
// to the console. If there is an error, the program
// will exit with a status code of 1.
func CheckError(err error) {
	// If there is an error
	if err != nil {
		// Print the error to the console
		fmt.Println(vars.ColorRed, "> error :", err, vars.ColorDefault)

		// Exit the program with a status code of 1
		os.Exit(1)
	}
}
