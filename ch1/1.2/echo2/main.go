// Echo2 prints command lines arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Notes:
// - os.Args[1:] has n omitted so it defaults to 0 or len(s).
// - Range produces a pair of values each iteration, the index and value of
// each element at that index.
// - Don't need an index so we just use a blank identifer (_) to satisfy the
// range syntax.
// - Garbage collection for s on each loop is expensive to looks to use Join.
