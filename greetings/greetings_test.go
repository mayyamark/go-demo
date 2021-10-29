package greetings

import (
	"regexp"
	"testing"
)

/*
	Naming convention: Test<smth-meaningful-about-the-test>
	Parameter: pointer to the testing package; it's used for logging/reporting from the test
	Execution: go test		-v: get verbose output (lists all of the tests and their results)
*/

// TestHelloName calls greetings.Hello with a name, checking for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Pesho"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmptyName(t *testing.T) {
	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
