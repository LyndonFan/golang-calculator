package main

import (
	"fmt"
	"math"
)

type Operator struct {
	Name string
	Arity int
	IsInfix bool
	Function func([]float64) (float64, error)
}

func getOperations() map[string]Operator{
	res := make(map[string]Operator)
	for name, op := range getUnaryOperations() {
		res[name] = op
	}
	for name, op := range getBinaryOperations() {
		res[name] = op
	}
	return res
}

func getUnaryOperations() map[string]Operator{
	originalFunctions := map[string]func(float64) (float64, error){
		"abs": func(f float64) (float64, error) {return math.Abs(f), nil},
		"exp": func(x float64) (float64, error) {return math.Exp(x), nil},
		"ln": saveLn,
	}
	res := make(map[string]Operator, len(originalFunctions))
	for name, f := range originalFunctions {
		res[name] = Operator{
			Name: name,
			Arity: 1,
			IsInfix: false,
			Function: func(args []float64) (float64, error) {
				if len(args) != 1 {
					return 0, fmt.Errorf("Invalid number of arguments")
				}
				return f(args[0])
			},
		}
	}
	return res
}

func getBinaryOperations() map[string]Operator{
	infixFunctions := map[string]func(float64, float64) (float64, error){
		"+": func(x, y float64) (float64, error) {return x+y, nil},
		"-": func(x, y float64) (float64, error) {return x-y, nil},
		"*": func(x, y float64) (float64, error) {return x*y, nil},
		"/": safeDivision,
		"%": cleverModulo,
		"^": pow,
	}
	namedFunctions := map[string]func(float64, float64) (float64, error){
		"max": max,
		"min": min,
		"pow": pow,
	}
	res := make(map[string]Operator, len(infixFunctions)+len(namedFunctions))
	for name, f := range infixFunctions {
		res[name] = Operator{
			Name: name,
			Arity: 2,
			IsInfix: true,
			Function: func(args []float64) (float64, error) {
				if len(args) != 2 {
					return 0, fmt.Errorf("Invalid number of arguments")
				}
				return f(args[0], args[1])
			},
		}
	}
	for name, f := range namedFunctions {
		res[name] = Operator{
			Name: name,
			Arity: 2,
			IsInfix: false,
			Function: func(args []float64) (float64, error) {
				if len(args) != 2 {
					return 0, fmt.Errorf("Invalid number of arguments")
				}
				return f(args[0], args[1])
			},
		}
	}
	return res
}

func saveLn (x float64) (float64, error) {
	if x <= 0 {
		return 0, fmt.Errorf("Invalid number: %f", x)
	}
	return math.Log(x), nil
}

func safeDivision(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Division by zero")
	}
	return a/b, nil
}

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

func max(a, b float64) (float64, error) {
	if a > b {
		return a, nil
	}
	return b, nil
}

func min(a, b float64) (float64, error) {
	if a < b {
		return a, nil
	}
	return b, nil
}

func pow(a, b float64) (float64, error) {
	return math.Pow(a, b), nil
}