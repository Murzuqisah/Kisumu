package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"kisumu/cmd/repl"
)

func main() {
	// bytes, _ := os.ReadFile("/examples/example.ksm")
	// source := string(bytes)

	// fmt.Printf("Code: %s", source)

	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user: %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s/%s\n", user.Username, runtime.GOOS)
	fmt.Println("Kisumu version: 0.1.0")

	fmt.Println("\nWelcome to the Kisumu programming language!")
	fmt.Println("Type 'help' for a list of available commands.")
	fmt.Printf("Feel free to type in commands, and let's start coding!\n\n")
	repl.Start(os.Stdin, os.Stdout)
}
