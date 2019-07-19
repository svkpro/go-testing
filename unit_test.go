package main

import (
	"testing"
)

/*
	Package testing provides support for automated testing of Go packages.
	It is intended to be used in concert with the “go test” command,
	which automates execution of any function of the form

		func TestXxx(*testing.T)

	where Xxx does not start with a lowercase letter.
*/

func fib(n int) [] int {
	if n <= 0 {
		return []int{0}
	}

	var fs []int
	fs = append(fs, 0)
	for i := 1; i <= n; i++ {
		if i == 1 {
			fs = append(fs, 1)
			continue
		}
		fs = append(fs, fs[i-1]+fs[i-2])
	}

	return fs
}

func TestCompareLengths(t *testing.T) {
	fst := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	fs := fib(10)
	if len(fst) != len(fs) {
		t.Error("Sequences have different lengths!")
	}
}

func TestCompareElements(t *testing.T) {
	fst := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	fs := fib(10)
	for i, v := range fst {
		if v != fs[i] {
			t.Error("Sequences are not match!")
		}
	}
}

/*
	go test -run ''      # Run all tests.
	go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
	go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
	go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
*/