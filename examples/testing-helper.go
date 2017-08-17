package main

import (
	"errors"
	"regexp"
	"testing"
)

//START,OMIT
func Something() error { return errors.New("oops") }

func TestSomething(t *testing.T) {
	err := Something()
	checkErr(t, err) // line 16
}

func checkErr(t *testing.T, err error) {
	// t.Helper()
	if err != nil {
		t.Fatal(err) // line 22
	}
}

// END,OMIT

var tests = []testing.InternalTest{
	{"TestSomething", TestSomething},
}

var benchmarks = []testing.InternalBenchmark{}

var examples = []testing.InternalExample{}

var matchPat string
var matchRe *regexp.Regexp

func matchString(pat, str string) (result bool, err error) {
	if matchRe == nil || matchPat != pat {
		matchPat = pat
		matchRe, err = regexp.Compile(matchPat)
		if err != nil {
			return
		}
	}
	return matchRe.MatchString(str), nil
}

func main() {
	testing.Main(matchString, tests, benchmarks, examples)
}
