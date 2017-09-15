// +build OMIT

package main

import "fmt"

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go func() {
		c <- sum(s[:len(s)/2]) // send sum to c
	}()
	go func() {
		c <- sum(s[len(s)/2:]) // send sum to c
	}()
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
