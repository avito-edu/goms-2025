package main

import (
	"errors"
	"fmt"
)

var (
	ErrDummy    = errors.New("damn")
	ErrNotDummy = errors.New("zamn")
)

func fail() error {
	return ErrDummy
}

func wrapper() error {
	err := fail()
	if err != nil {
		return fmt.Errorf("err wrpapped %w", err)
	}

	return nil
}

func main() {
	errZamn := fail()
	if errZamn != nil {
		fmt.Println("handle error")
	}

	err := wrapper()
	//  whether any error in err's tree matches target.
	if errors.Is(err, ErrDummy) {
		fmt.Println("got err that matches target")
	}

	if errors.Is(err, ErrNotDummy) {
		fmt.Println("wottahell")
	}
}
