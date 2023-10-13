package main

import (
	"fmt"
	"strings"
)

func isDigit(b byte) bool {
	return byte('0') <= b && b <= byte('9')
}

func getVariableName(tokens []string) (string, []string, error) {
	if len(tokens) < 2 || tokens[1] != "=" {
		return "", tokens, nil
	}
	variableName := tokens[0]
	err := CheckValidCacheKey(variableName)
	if err != nil {
		return "", []string{}, err
	}
	operations := getOperations()
	_, exists := operations[variableName]
	if exists {
		return "", []string{}, fmt.Errorf("Variable name %s already exists", variableName)
	}
	return variableName, tokens[2:], nil
}
func genIsOperator(operations map[string]Operator) func(byte) bool {
	singleSymbolOperators := make(map[string]bool, len(operations))
	for s := range operations {
		if len(s) == 1 {
			singleSymbolOperators[s] = true
		}
	}
	for _, b := range []byte("(),=") {
		singleSymbolOperators[string(b)] = true
	}
	return func(s byte) bool {
		_, exists := singleSymbolOperators[string(s)]
		return exists
	}
}

func tokenize(s string, cache *Cache) ([]string, error) {
	operations := getOperations()
	isOperator := genIsOperator(operations)

	s = strings.ReplaceAll(s, " ", "")
	tokens := make([]string, 0, len(s))
	currToken := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		if isDigit(b) || b == '.' {
			// extend number
			// log.Printf("extending number with %c\n", b)
			if len(currToken) > 0 {
				lastSymbol := currToken[len(currToken)-1]
				shouldExtend := isDigit(lastSymbol) || lastSymbol == '.'
				if lastSymbol == '-' {
					shouldExtend = true
					if len(tokens) > 0 {
						prevToken := tokens[len(tokens)-1]
						shouldExtend = !representsNumber(prevToken, cache)
					}
				}
				// log.Printf("shouldExtend = %t\n", shouldExtend)
				if !shouldExtend {
					tokens = append(tokens, string(currToken))
					currToken = currToken[:0]
				}
			}
			currToken = append(currToken, b)
		} else if isAlpha(string(b)) {
			// extend alpha
			// log.Printf("extending alpha with %c\n", b)
			if len(currToken) > 0 {
				shouldExtend := isAlpha(string(currToken))
				if !shouldExtend {
					tokens = append(tokens, string(currToken))
					currToken = currToken[:0]
				}
			}
			currToken = append(currToken, b)
		} else if !isOperator(b) && b != '=' {
			err := fmt.Errorf("invalid token: %c", b)
			return nil, err
		} else if b == '-' {
			// handle minus sign
			// log.Printf("handling minus sign with %c\n", b)
			if len(currToken) > 0 {
				tokens = append(tokens, string(currToken))
				currToken = currToken[:0]
			}
			currToken = append(currToken, b)
		} else {
			// other operators
			// log.Printf("handling operator with %c\n", b)
			if len(currToken) > 0 {
				tokens = append(tokens, string(currToken))
				currToken = currToken[:0]
			}
			currToken = append(currToken, b)
		}
		// log.Printf("currToken = %s, tokens = %s\n", currToken, tokens)
	}
	if len(currToken) > 0 {
		tokens = append(tokens, string(currToken))
	}
	return tokens, nil
}
