// Prints command line arguments.
package main

import (
	"os"
	"strings"
)

func ArgConcat() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
}

func ArgRange() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
}

func ArgJoin() {
	strings.Join(os.Args[1:], " ")
}
