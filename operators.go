package main

import (
	"fmt"
	"math"
)

func cleverModulo(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Modulo by zero")
	}
	if b<0 {
		res, err := cleverModulo(-a, -b)
		return -res, err
	}
	d := a/b
	return a-math.Floor(d)*b, nil
}