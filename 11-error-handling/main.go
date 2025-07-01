package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("no division by zero allowed")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 5)
	if err == nil {
		fmt.Println("Division result is", result)
	} else {
		fmt.Println(err)
	}

	result2, err2 := divide(10, 0)
	if err2 == nil {
		fmt.Println("Division result is", result2)
	} else {
		fmt.Println(err2)
	}
}
