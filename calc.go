package calc

import (
	"errors"
	"fmt"
)

var ErrNotDivisibleByZero = errors.New("not divisible by zero")

func Add(a, b float64) float64 {
	return a + b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("calc.Divide: provided %d for denominator %w", b, ErrNotDivisibleByZero)
	}

	return a / b, nil
}
