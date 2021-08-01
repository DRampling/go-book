// Exercise 1.2:
// Modify the echo program to print the index and value for each of its
// arguments, one per line.

// Prints the index and value of its arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
