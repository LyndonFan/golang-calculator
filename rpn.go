package main

import (
	"fmt"
	"strconv"
)

func evaluateRPN(tokens []string, cache *Cache) (float64, error) {
	if len(tokens) == 0 {
		return 0, fmt.Errorf("Empty stack")
	}
	stack := make([]float64, 0)
	constants := getConstants()
	operations := getOperations()
	var err error
	for _, token := range tokens {
		op, opExists := operations[token]
		if opExists {
			args := make([]float64, op.Arity)
			for i := 0; i < op.Arity; i++ {
				args[i] = stack[len(stack)-1-i]
			}
			number, err := op.Function(args)
			if err != nil {
				return 0, err
			}
			stack = stack[:len(stack)-op.Arity]
			stack = append(stack, number)
		} else {
			number, varExists := cache.Get(token)
			if !varExists {
				number, varExists = constants[token]
				if !varExists {
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