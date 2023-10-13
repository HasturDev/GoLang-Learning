// Go program to illustrate the
// concept of init() function

// Declaration of the main package
package main

// Importing package
import (
	"fmt"

	statistics "xithulhu.com/Statistics"
)

// Multiple init() function

// items that we'll use later for dictionaries or json or somethign
func init() {
	fmt.Println("Welcome to init() function")
}

// this init is going to initialize items that we'll use later in structs
func init() {
	fmt.Println("Hello! init() function")
}

// Main function
func main() {
	// func (r *Repository) Log(o *LogOptions) (object.CommitIter, error)
	fmt.Println("Welcome to main() function")
	myCode := statistics.Code{FileSomething: "File1", CodeSomething: "Some Code"}

	result := statistics.Search(myCode)
	fmt.Println(result)
}
