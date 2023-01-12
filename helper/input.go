package helper

import (
	"bufio"
	"errors"
	"fmt"
	"go-supabase/handler"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var scanner = bufio.NewScanner(os.Stdin)

// getUserInput prompts the user with the provided
// question and returns their response as a string.
func GetUserInput(question string) string {
	// Print the question to the console
	fmt.Print("\n"+question)

	// scanner is Global variable
	// Attempt to read the user's response
	if !scanner.Scan() {
		// If there was an error reading the response,
		// pass an error to the CheckError function
		handler.CheckError(errors.New("scanning error"))
	}

	// Return the user's response
	return scanner.Text()
}

// safe-action for destroy/ hardreset
func VerifyUserAction() bool {
	// Initialize variables
	var randomInt int
	var userInput uint64
	var err error

	// Generate a random integer 
	// between 1000 and 9999
	rand.Seed(time.Now().Unix())
	randomInt = rand.Intn(9999-1000) + 1000

	// Prompt the user to re-type the pin
	fmt.Println("Verify your action !\nRe-type the pin : ", randomInt)

	// Get user input and convert it to a uint64
	userInput, err = strconv.ParseUint(GetUserInput("Pin : "), 10, 64)
	handler.CheckError(err)

	// Return whether the user input 
	// matches the random integer
	return randomInt == int(userInput)
}