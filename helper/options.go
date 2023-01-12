package helper

import "fmt"

// prints a list of options to the console.
func ShowOptions() {
	// Print a header message
	fmt.Println("\nOptions :")
	fmt.Println("0. Exit App					6. Update Todo's Status")
	fmt.Println("1. Show Todos					7. SoftDelete Todo")
	fmt.Println("2. Show All Todos (W/ Trashed Todos)		8. Restore Todo (From SoftDelete)")
	fmt.Println("3. Show Just Trashed Todos			9. Destroy (Delete Permanent)")
	fmt.Println("4. Create New Todo				10. Hard Reset (Drop -> re-create Table)")
	fmt.Println("5. Edit Todo")
}