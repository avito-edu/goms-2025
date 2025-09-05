package main

import (
	"errors"
	"fmt"
)

type ErrorDummy struct { /* ... */
}

// Dummy — implements error interface
func (dm *ErrorDummy) Error() string {
	return "damn Daniel"
}

func fais() error {
	return new(ErrorDummy)
}

var (
	ErrUsingErrorsPkg = errors.New("")
	ErrUsingFmtPkg    = fmt.Errorf("you %w", ErrUsingErrorsPkg)
)

func main() {
	if err := fais(); err != nil {
		fmt.Printf("got: %+w", err)
	}
}
