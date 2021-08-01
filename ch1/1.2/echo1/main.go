// It's considered good practice to describe each package before declaration
package main

import (
	"fmt"
	"os"
)

func main() {
	// Not explicitly initialized - implicitly set to zero value for zero type
	var s, sep string
	for i := 1; 1 < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
