package main

import (
	"fmt"
	"math"
	"strconv"
)

func evaluateRPN(tokens []string, cache *Cache) (float64, error) {
	if len(tokens) == 0 {
		return 0, fmt.Errorf("Empty stack")
	}
	stack := make([]float64, 0)
	constants := getConstants()
	var number float64
	var err error
	var exists bool
	for _, token := range tokens {
		switch token {
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2] + stack[len(stack)-1])
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2] - stack[len(stack)-1])
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack)-2] * stack[len(stack)-1])
		case "/":
			if stack[len(stack)-1] == 0.0 {
				return 0, fmt.Errorf("Division by zero")
			}
			stack = append(stack[:len(stack)-2], stack[len(stack)-2] / stack[len(stack)-1])
		case "%":
			remainder, err := cleverModulo(stack[len(stack)-2], stack[len(stack)-1])
			if err != nil {
				return 0, err
			}
			stack = append(stack[:len(stack)-2], remainder)
		case "^":
			stack = append(stack[:len(stack)-2], math.Pow(stack[len(stack)-2], stack[len(stack)-1]))
		default:
			number, exists = cache.Get(token)
			if !exists {
				number, exists = constants[token]
				if !exists {
					number, err = strconv.ParseFloat(token, 64)
				}
			}
			if err != nil {
				return 0, err
			}
			stack = append(stack, number)
		}
	}
	return stack[0], nil
}