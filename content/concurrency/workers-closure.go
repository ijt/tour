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

var bigData = []bool{
	false, false, false, false,
	true, false, true, false,
	true, true, true, true,
}

func main() {
	// Start the workers.
	ch := make(chan result)
	for i := 0; i < numWorkers; i++ {
		start, end := workerRange(i)
		go func(id int, bits []bool) {
			n, err := countTrue(bits)
			ch <- result{id: id, n: n, err: err}
		}(i+1, bigData[start:end])
	}

	// Wait for the workers to finish.
	n := 0
	for i := 0; i < numWorkers; i++ {
		r := <-ch
		if r.err != nil {
			fmt.Printf("worker #%d failed: %v\n", r.id, r.err)
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
