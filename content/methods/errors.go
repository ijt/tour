// +build OMIT

package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	fmt.Printf("err = %v\n", err)
	err = errors.New("failed to launch")
	fmt.Printf("err.Error() = %q\n", err.Error())
	fmt.Printf("err = %v\n", err)
}
