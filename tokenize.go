package main

import (
	"fmt"
	"strings"
)

func isDigit(b byte) bool{
	// 48 = byte('0'), 57 = byte('9')
	return 48<=b && b<=57
}

func isOperator(s string) bool{
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
		"^": true,
		"%": true,
		"(": true,
		")": true,
	}
	return operators[s]
}

func getVariableName(tokens []string) (string, []string, error) {
	if len(tokens) < 2 || tokens[1] != "=" {return "", tokens, nil}
	variableName := tokens[0]
	err := CheckValidCacheKey(variableName)
	if err != nil {
		return "", []string{}, err
	}
	return variableName, tokens[2:], nil
}

func tokenize(s string, cache *Cache) ([]string, error) {
	s = strings.ReplaceAll(s, " ", "")
	tokens := make([]string, 0, len(s))
	currToken := make([]byte, 0, len(s))
	for _, b := range([]byte(s)){
		if isDigit(b) || b == '.' {
			if len(currToken) > 0 {
				shouldExtend := isDigit(currToken[len(currToken)-1])
				shouldExtend = shouldExtend || currToken[len(currToken)-1] == '.'
				if currToken[len(currToken)-1] == '-' {
					shouldExtend = true
					if len(tokens) > 0 {
						prevToken := tokens[len(tokens)-1]
						shouldExtend = !representsNumber(prevToken, cache)
					}
				}
				if !shouldExtend {
					tokens = append(tokens, string(currToken))
					currToken = currToken[:0]
				}
			}
			currToken = append(currToken, b)
		} else if isAlpha(string(b)) {
			if len(currToken) > 0 {
				shouldExtend := isAlpha(string(currToken))
				if !shouldExtend {
					tokens = append(tokens, string(currToken))
					currToken = currToken[:0]
				}
			}
			currToken = append(currToken, b)
		} else if !isOperator(string(b)) && b != '=' {
			err := fmt.Errorf("invalid token: %c", b)
			return nil, err	
		} else if b == '-' {
			if len(currToken) > 0 {
				tokens = append(tokens, string(currToken))
				currToken = currToken[:0]
			}
			currToken = append(currToken, b)
		} else { // other operators
			if len(currToken) > 0 {
				tokens = append(tokens, string(currToken))
				currToken = currToken[:0]
			}
			tokens = append(tokens, string(b))
		}
	}
	if len(currToken) > 0 {
		tokens = append(tokens, string(currToken))
	}
	return tokens, nil
}