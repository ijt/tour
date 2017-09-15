// +build OMIT

package main

import (
	"errors"
	"fmt"
)

// Greet prints a greeting for the given subject to stdout.
func Greet(subject string) error {
	// Check for invalid inputs.
	if subject == "" {
		return errors.New("no subject to greet")
	}
	if len(subject) > 5 {
		return fmt.Errorf("greet %s...: subject too large", subject[:5])
	}

	// Other functions can return errors, too.
	// Here we add context about the operation that failed to the underlying error.
	if _, err := fmt.Println("Hello,", subject); err != nil {
		return fmt.Errorf("greet %s: %v", subject, err)
	}
	return nil
}

func main() {
	if err := Greet("world"); err != nil {
		fmt.Println(err)
	}
}
