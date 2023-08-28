package main

import (
	"fmt"
	"os"
)

type Command func() error

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": Exit,
		"quit": Exit,
		"help": Help,
	}
}

func Exit() error {
	os.Exit(0)
	return nil
}

func Help() error {
	fmt.Println("This is a CLI calculator")
	fmt.Println("Enter 'help' to get this message")
	fmt.Println("Enter 'exit' or 'quit' to exit")
	fmt.Println("Enter anything else to calculate")
	return nil
}