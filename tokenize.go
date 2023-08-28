package main

import (
	"fmt"
	"log"
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

func tokenize(s string) ([]string, error) {
	s = strings.ReplaceAll(s, " ", "")
	log.Println(fmt.Sprintf("tokenize: \"%s\"", s))
	tokens := make([]string, 0, len(s))
	currToken := make([]byte, 0, len(s))
	for _, b := range([]byte(s)){
		if isDigit(b) || b == '.' {
			if len(currToken) > 0 && currToken[len(currToken)-1] == '-' {
				if len(tokens) > 0 {
					prevToken := tokens[len(tokens)-1]
					if isDigit(prevToken[len(prevToken)-1]) {
						tokens = append(tokens, string(currToken))
						currToken = currToken[:0]
					}
				}
			}
			currToken = append(currToken, b)
		} else if !isOperator(string(b)) {
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