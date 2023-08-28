package main

import (
	"fmt"
	"math"
	"strconv"
)

func evaluateRPN(tokens []string) (float64, error) {
	if len(tokens) == 0 {
		return 0, fmt.Errorf("Empty stack")
	}
	stack := make([]float64, 0)
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
			number, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, number)
		}
	}
	return stack[0], nil
}