package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func main() {
	// using %v
	err1 := fmt.Errorf("wrap with v: %v", ErrNotFound)

	// using %w
	err2 := fmt.Errorf("wrap with w: %w", ErrNotFound)

	fmt.Println("errors.Is(err1, ErrNotFound):", errors.Is(err1, ErrNotFound))
	fmt.Println("errors.Is(err2, ErrNotFound):", errors.Is(err2, ErrNotFound))
	fmt.Println(err1)
	fmt.Println(err2)
}
