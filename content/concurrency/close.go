// +build OMIT

package main

import (
	"fmt"
)

func count(n int, c chan int) {
	defer close(c)
	for i := 0; i < n; i++ {
		c <- i
	}
}

func main() {
	c := make(chan int)
	go count(10, c)
	for {
		i, ok := <-c
		if !ok {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("receive closed:", <-c)
}
