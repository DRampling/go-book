// Exercise 1.3:
// Experiment to measure the difference in running time between our potentially
// inefficient versions and one that uses strings.Join.

// go test -bench=. -args a b c d e

// Compares string concatenation
package main

import (
	"testing"
)

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArgConcat()
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArgRange()
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ArgJoin()
	}
}
