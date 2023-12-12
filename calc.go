package calc

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

var ErrNotDivisibleByZero = errors.New("not divisible by zero")

func Add(a, b float64) (float64, color.Attribute) {
	return a + b, color.FgGreen
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("calc.Divide: provided '%d' for denominator %w", b, ErrNotDivisibleByZero)
	}

	return a / b, nil
}

func Sub(a, b int) int {
	return a - b
}

func Exp(a, b int) int {
	return a ^ b
}
