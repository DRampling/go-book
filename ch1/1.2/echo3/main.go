// Echo4 prints command lines arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Notes:
// - Option to print a slice using fmt.Println(os.Args[1:]) for debugging.
