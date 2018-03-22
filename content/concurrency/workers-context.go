// +build OMIT

package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func countTrue(ctx context.Context, bits []bool) (int, error) {
	n := 0
	for _, b := range bits {
		if b {
			n++
			select {
			case <-time.After(1 * time.Second):
			case <-ctx.Done():
				return 0, ctx.Err()
			}
		}
	}
	if n == 0 {
		return 0, errors.New("nothing is true!")
	}
	return n, nil
}

const numWorkers = 3

var bigData = []bool{
	false, false, false, false,
	true, false, true, false,
	true, true, true, true,
}

func main() {
	// Start the workers.
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan result)
	for i := 0; i < numWorkers; i++ {
		start, end := workerRange(i)
		go func(id int, bits []bool) {
			n, err := countTrue(ctx, bits)
			ch <- result{id: id, n: n, err: err}
		}(i+1, bigData[start:end])
	}

	// Wait for the workers to finish.
	// Stop on first failure.
	n := 0
	for i := 0; i < numWorkers; i++ {
		r := <-ch
		if r.err != nil {
			fmt.Printf("worker #%d failed: %v\n", r.id, r.err)
			cancel() // cancel no-ops after first call
			continue
		}
		fmt.Printf("worker #%d reported %d\n", r.id, r.n)
		n += r.n
	}
	fmt.Printf("total: %d\n", n)
}

type result struct {
	id  int
	n   int
	err error
}

func workerRange(i int) (start, end int) {
	itemsPerWorker := (len(bigData) + numWorkers - 1) / numWorkers
	end = (i + 1) * itemsPerWorker
	if end > len(bigData) {
		end = len(bigData)
	}
	return i * itemsPerWorker, end
}
