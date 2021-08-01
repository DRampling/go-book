// It's considered good practice to describe each package before declaration
package main

import (
	"fmt"
	"os"
)

func main() {
	// Not explicitly initialized - implicitly set to zero value for zero type
	var s, sep string
	// Short variable declartion assigns type baseds on initialized values
	for i := 1; 1 < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Traditional for loop:
//  for initialization; boolean expression; post {
//      zero or more statements
//  }

// Traditional while loop:
//  for boolean expression {
//      zero or more statements
//  }

// Traditional infinate loop:
//  for {
//      break or return to escape
//  }
