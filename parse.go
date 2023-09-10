package main

import (
	"fmt"
	"strings"
)

func getPrecendenceLevel(token string) int {
	if token == "+" || token == "-" {
		return 1
	} else if token == "*" || token == "/" {
		return 2
	} else if token == "^" {
		return 3
	} else if token == "%" {
		return 4
	}
	return 0
}

func isRightAssociative(token string) bool {
	return token == "^"
}

func lastOperatorExistsAndIsnotParen(operatorStack []string) bool{
	return len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "("
}

// https://en.wikipedia.org/wiki/Shunting_yard_algorithm#The_algorithm_in_detail

func convertToRPN(tokens []string, cache *Cache) ([]string, error ){
	result := make([]string, 0, len(tokens))
	operatorStack := make([]string, 0, len(tokens))
	lastWasOperator := false
	for i, token := range(tokens){
		if strings.Count(token, ".") > 1 {
			err := fmt.Errorf("Invalid token with more than 1 decimal point: %s", token)
			return []string{}, err
		}
		if token == "." || token == "-." {
			err := fmt.Errorf("Invalid token: %s", token)
			return []string{}, err
		}
		if representsNumber(token, cache){
			if !lastWasOperator && i>0 {
				err := fmt.Errorf("Consecutive numbers found: %s then %s", tokens[i-1], token)
				return []string{}, err
			}
			result = append(result, strings.ToLower(token))
			lastWasOperator = false
			continue
		}
		if token == "("{
			operatorStack = append(operatorStack, token)
			lastWasOperator = true
		} else if token == ")" {
			lastWasOperator = true
			for lastOperatorExistsAndIsnotParen(operatorStack){
				result = append(result, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			if len(operatorStack) == 0{
				return []string{}, fmt.Errorf("Unmatched )")
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else if !isOperator(token) {
			err := fmt.Errorf("Invalid token: %s", token)
			return []string{}, err
		} else {
			lastWasOperator = true
			for lastOperatorExistsAndIsnotParen(operatorStack){
				lastOperator := operatorStack[len(operatorStack)-1]
				lastLevel := getPrecendenceLevel(lastOperator)
				currentLevel := getPrecendenceLevel(token)
				if lastLevel < currentLevel{
					break
				}
				if lastLevel == currentLevel && isRightAssociative(token){
					break
				}
				result = append(result, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, token)
		}
	}
	for len(operatorStack) > 0 {
		op := operatorStack[len(operatorStack)-1]
		if op == "("{
			return []string{}, fmt.Errorf("Unmatched (")
		}
		result = append(result, op)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}
	return result, nil
}