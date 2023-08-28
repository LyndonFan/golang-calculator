package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculate(input string, cache *Cache) (float64, error){
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
				fmt.Printf("%s = %f\n",input,res)
			}
		}
	}
}
