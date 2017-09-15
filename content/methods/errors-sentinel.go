// +build OMIT

package main

import (
	"errors"
	"fmt"
)

// ErrMissingSubject is returned when calling Greet("").
var ErrMissingSubject = errors.New("no subject to greet")

// Greet prints a greeting for the given subject to stdout.
// If the subject is an empty string, then Greet returns
// ErrMissingSubject.
func Greet(subject string) error {
	if subject == "" {
		return ErrMissingSubject
	}
	if len(subject) > 5 {
		return fmt.Errorf("greet %s...: subject too large", subject[:5])
	}
	if _, err := fmt.Println("Hello,", subject); err != nil {
		return fmt.Errorf("greet %s: %v", subject, err)
	}
	return nil
}

func main() {
	// We can take different action based on the kind of error.
	if err := Greet(""); err == ErrMissingSubject {
		fmt.Println("Forgot who to talk to!")
	} else if err != nil {
		fmt.Println(err)
	}
}
