package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// the main function is a cli app
// that takes in the user's input
// if it's "exit", runs os.exit
// otherwise runs the tokenize function
// and prints the result

func getHelp() error {
	fmt.Println("This is a CLI calculator")
	fmt.Println("Enter 'help' to get this message")
	fmt.Println("Enter 'exit' or 'quit' to exit")
	fmt.Println("Enter anything else to calculate")
	return nil
}

func calculate(input string) (float64, error){
	tokens, err := tokenize(input)
	if err != nil {
		return 0, err
	}
	rpn, err := convertToRPN(tokens)
	if err != nil {
		return 0, err
	}
	log.Println(rpn)
	return evaluateRPN(rpn)
}


func main() {
	_ = getHelp()
	reader := bufio.NewReader(os.Stdin)
	var input string
	for true {
		fmt.Print("> ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		switch input {
			case "exit":
				os.Exit(0)
			case "quit":
				os.Exit(0)
			case "help":
				getHelp()
			default:
				res, err := calculate(input)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("%s = %f\n",input,res)
				}
		}
	}
}
