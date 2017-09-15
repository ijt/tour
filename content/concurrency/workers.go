// +build OMIT

package main

import "fmt"

func run() error {
	return nil
}

func main() {
	errc := make(chan error)
	// Start the workers.
	for i := 0; i < 10; i++ {
		go func() {
			errc <- run()
		}()
	}
	// Wait for the workers to finish.
	for i := 0; i < 10; i++ {
		if err := <-errc; err != nil {
			fmt.Println("failed a run:", err)
		}
	}
}
