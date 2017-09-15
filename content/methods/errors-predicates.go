// +build OMIT

package main

import (
	"errors"
	"fmt"
)

// errMissingSubject is returned when calling Greet("").
var errMissingSubject = errors.New("no subject to greet")

// largeSubjectError is returned when Greet receives an input that is too large.
// It holds the leading prefix of the input.
type largeSubjectError string

// Error implements the error interface.
func (e largeSubjectError) Error() string {
	return "greet " + string(e) + "...: subject too large"
}

// Greet prints a greeting for the given subject to stdout.
// If the subject is an empty string, then Greet returns
// ErrMissingSubject.
func Greet(subject string) error {
	if subject == "" {
		return errMissingSubject
	}
	if len(subject) > 5 {
		return largeSubjectError(subject[:5])
	}
	if _, err := fmt.Println("Hello,", subject); err != nil {
		return fmt.Errorf("greet %s: %v", subject, err)
	}
	return nil
}

// IsMissingSubject reports whether e indicates a failure to provide a
// greeting subject.
func IsMissingSubject(e error) bool {
	return e == errMissingSubject
}

// IsLargeSubject reports whether e indicates that a large input was
// given to Greet.
func IsLargeSubject(e error) bool {
	_, ok := e.(largeSubjectError)
	return ok
}

func main() {
	// We can take different action based on the kind of error.
	if err := Greet(""); IsMissingSubject(err) {
		fmt.Println("Forgot who to talk to!")
	} else if IsLargeSubject(err) {
		fmt.Println("Insufficient memory")
	} else if err != nil {
		fmt.Println(err)
	}
}
