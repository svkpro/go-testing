package main

import (
	"testing"
)

/*
	The Go testing package contains a benchmarking facility that can be used to examine the performance of your Go code.
*/

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func BenchmarkFactorial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Factorial(10)
	}
}

/*
	Run go test -bench=. to run benchmark test
*/