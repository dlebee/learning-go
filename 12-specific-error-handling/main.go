package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrDivideByZero = errors.New("division by zero")
	ErrOverflow     = errors.New("integer overflow")
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}

	// simulate overflow case — in reality Go ints won't easily overflow
	if a == math.MaxInt32 && b == 1 {
		return 0, ErrOverflow
	}

	return a / b, nil
}

func main() {
	cases := []struct {
		a, b int
	}{
		{10, 5},
		{10, 0},
		{math.MaxInt32, 1},
	}

	for _, c := range cases {
		result, err := divide(c.a, c.b)
		if err != nil {
			switch {
			case errors.Is(err, ErrDivideByZero):
				fmt.Printf("Cannot divide %d by %d: division by zero\n", c.a, c.b)
			case errors.Is(err, ErrOverflow):
				fmt.Printf("Overflow occurred during division: %d / %d\n", c.a, c.b)
			default:
				fmt.Println("Unknown error:", err)
			}
		} else {
			fmt.Printf("Result of %d / %d = %d\n", c.a, c.b, result)
		}
	}
}
