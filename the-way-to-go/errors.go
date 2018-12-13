package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}

