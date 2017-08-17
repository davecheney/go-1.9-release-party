package main

import (
	"errors"
	"testing"
)

//START,OMIT
func Something() error { return errors.New("oops") }

func TestSomething(t *testing.T) {
	err := Something()
	checkErr(t, err)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// END,OMIT

func main() {
	testing.RunTests(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{
			Name: "TestSomething",
			F:    TestSomething,
		}})
}
