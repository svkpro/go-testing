package main

import (
	"os"
	"testing"
)
/*
	It is sometimes necessary for a test program to do extra setup or teardown before or after testing.
	It is also sometimes necessary for a test to control which code runs on the main thread.
*/
func setup() {
	println("setup")
}

func teardown() {
	println("teardown")
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

/*
	go test -run ''      # Run all tests.
	go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
	go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
	go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
*/