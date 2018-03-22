// +build OMIT

package main

import "fmt"

func countTrue(bits []bool) (int, error) {
	n := 0
	for _, b := range bits {
		if b {
			n++
		}
	}
	return n, nil
}

const numWorkers = 3

func main() {
	// Start the workers.
	ch := make(chan result)
	for i := 0; i < numWorkers; i++ {
		go func() {
			n, err := countTrue([]bool{true, false, true, false})
			ch <- result{n: n, err: err}
		}()
	}

	// Wait for the workers to finish.
	n := 0
	for i := 0; i < numWorkers; i++ {
		r := <-ch
		if r.err != nil {
			fmt.Printf("worker failed: %v\n", r.err)
			continue
		}
		fmt.Printf("worker reported %d\n", r.n)
		n += r.n
	}
	fmt.Printf("total: %d\n", n)
}

type result struct {
	n   int
	err error
}
