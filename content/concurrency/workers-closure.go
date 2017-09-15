// +build OMIT

package main

import (
	"fmt"
	"time"
)

func run(d time.Duration) error {
	fmt.Println("running for", d)
	time.Sleep(d)
	return nil
}

func main() {
	errc := make(chan error)
	// Start the workers.
	for i := 0; i < 10; i++ {
		d := time.Duration(i) * time.Second
		go func() {
			// Try moving the assignment of d into the goroutine.
			errc <- run(d)
		}()
	}
	// Wait for the workers to finish.
	for i := 0; i < 10; i++ {
		if err := <-errc; err != nil {
			fmt.Println("failed a run:", err)
		}
	}
}
