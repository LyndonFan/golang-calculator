package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func calculate(input string) (float64, error){
	log.Printf("input:  \"%s\"", input)
	tokens, err := tokenize(input)
	if err != nil {
		return 0, err
	}
	log.Printf("tokens: %s", tokens)
	rpn, err := convertToRPN(tokens)
	if err != nil {
		return 0, err
	}
	log.Printf("rpn:    %s", rpn)
	return evaluateRPN(rpn)
}


func main() {
	commands := getCommands()
	commands["help"]()
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
			res, err := calculate(input)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%s = %f\n",input,res)
			}
		}
	}
}
