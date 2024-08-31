package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("/examples/00.ksm")
	source := string(bytes)

	fmt.Printf("Code: %s", source)
}