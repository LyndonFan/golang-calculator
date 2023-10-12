package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculate(input string, cache *Cache) (float64, error) {
	log.Printf("input:  \"%s\"", input)
	tokens, err := tokenize(input, cache)
	if err != nil {
		return 0, err
	}
	variableName, tokens, err := getVariableName(tokens)
	if err != nil {
		return 0, err
	}
	log.Printf("tokens: %s", tokens)
	rpn, err := convertToRPN(tokens, cache)
	if err != nil {
		return 0, err
	}
	log.Printf("rpn:    %s", rpn)
	res, err := evaluateRPN(rpn, cache)
	if err != nil {
		return 0, err
	}
	if variableName != "" {
		cache.Set(variableName, res)
	}
	return res, nil
}

func main() {
	commands := getCommands()
	commands["help"]()
	cache := NewCache()
	reader := bufio.NewReader(os.Stdin)
	operations := getOperations()
	for s, op := range operations {
		fmt.Printf("%s: %s, ", s, op.Name)
		if op.Arity == 1 {
			res, _ := op.Function([]float64{3})
			fmt.Printf("e.g. %s(3) = %.2f\n", op.Name, res)
		} else if op.Arity == 2 {
			res, _ := op.Function([]float64{2, 5})
			if op.IsInfix {
				fmt.Printf("e.g. 2 %s 5 = %.2f\n", op.Name, res)
			} else {
				fmt.Printf("e.g. %s(2, 5) = %.2f\n", op.Name, res)
			}
		}
	}
	var input string
	for true {
		fmt.Print("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		command, exists := commands[input]
		if exists {
			err := command()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			res, err := calculate(input, cache)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%s = %f\n", input, res)
			}
		}
	}
}
