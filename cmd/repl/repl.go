package repl

import (
	"bufio"
	"fmt"
	"io"

	"kisumu/pkg/lexer"
)

const PROMPT = "kisumu $"

// Start is a read-eval-print loop (REPL) for the Kisumu programming language.
// It continuously reads input from the provided reader, tokenizes it,
// and writes the token type and literal to the provided writer.
func Start(in io.Reader, out io.Writer) {
	// scanner := bufio.NewScanner(in)

	// for {
	// 	fmt.Printf(PROMPT)
	// 	scanned := scanner.Scan()
	// 	if !scanned {
	// 		return
	// 	}

	// 	line := scanner.Text()
	// 	lexer := lexer.Tokenize(line)

	// 	for tok := lexer.GetNextToken(); tok.Type != lexer.EOF; tok = lexer.GetNextToken() {
	// 		fmt.Printf("%v\n", tok.Type, tok.Literal)
	// 	}

	// }

	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for {
		fmt.Fprint(writer, PROMPT)
		writer.Flush() // Ensure prompt is printed before reading input

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(writer, "Error reading input: %v\n", err)
				writer.Flush()
				continue
			}
			break // Exit loop on EOF or error
		}

		line := scanner.Text()
		Lexer := lexer.Tokenize(line)

		for {
			tok := Lexer.GetNextToken()
			if tok.Type == lexer.EOF {
				break
			}
			if tok.Type == lexer.ILLEGAL {
				fmt.Fprintf(writer, "Illegal token: %s\n", tok.Literal)
			} else {
				fmt.Fprintf(writer, "Token: %s (%s)\n", tok.Type, tok.Literal)
			}
		}
		writer.Flush() // Ensure output is written
	}
}

type REPL struct {
	prompt string
	lexer  *lexer.Lexer
	// ksm    *lexer
	// lexer *lexer
	// env    *Environment
	// runner *runner
	input  *bufio.Reader
	output *bufio.Writer
}

func NewREPL(prompt string, lexer *lexer.Lexer, input *bufio.Reader, output *bufio.Writer) *REPL {
	return &REPL{
		prompt: prompt,
		lexer:  lexer,
		// ksm:    ksm,
		// lexer: lexer,
		// env:    env,
		// runner: runner,
		input:  input,
		output: output,
	}
}
