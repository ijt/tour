// +build OMIT

package main

import (
	"context"
	"fmt"
	"time"
)

func run(ctx context.Context, d time.Duration) error {
	fmt.Println("running for", d)
	defer fmt.Println("done")

	select {
	case <-time.After(d):

	// Context's cancelation signal.
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	errc := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	// Start the workers.
	for i := 0; i < 10; i++ {
		d := time.Duration(i) * time.Second
		go func() {
			errc <- run(ctx, d)
		}()
	}
	// Wait for the first worker to finish.
	if err := <-errc; err != nil {
		fmt.Println("failed first run:", err)
	}
	// Cancel the other workers, wait for them to finish.
	cancel()
	for i := 0; i < 9; i++ {
		if err := <-errc; err != nil {
			fmt.Println("failed a later run:", err)
		}
	}
}
