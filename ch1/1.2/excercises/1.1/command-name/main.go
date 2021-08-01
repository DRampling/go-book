// Exercise 1.1:
// Modify the echo statement to also print os.Args[0], the name of the command
// that invoked it.

// Prints the command that invoked it.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
