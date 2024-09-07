package main

import (
	"fmt"
	"os"
)

func main() {
	bytes, _ := os.ReadFile("/examples/example.ksm")
	source := string(bytes)

	fmt.Printf("Code: %s", source)
}
