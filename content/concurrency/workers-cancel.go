// +build OMIT

package main

import (
	"errors"
	"fmt"
	"time"
)

func run(cancel chan struct{}, d time.Duration) error {
	fmt.Println("running for", d)
	defer fmt.Println("done")

	select {
	case <-time.After(d):
	case <-cancel:
		return errors.New("canceled")
	}
	return nil
}

func main() {
	errc := make(chan error)
	cancel := make(chan struct{})
	// Start the workers.
	for i := 0; i < 10; i++ {
		d := time.Duration(i) * time.Second
		go func() {
			errc <- run(cancel, d)
		}()
	}
	// Wait for the first worker to finish.
	if err := <-errc; err != nil {
		fmt.Println("failed first run:", err)
	}
	// Cancel the other workers, wait for them to finish.
	close(cancel)
	for i := 0; i < 9; i++ {
		if err := <-errc; err != nil {
			fmt.Println("failed a later run:", err)
		}
	}
}
