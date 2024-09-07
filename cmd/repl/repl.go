package repl

import (
	"bufio"
	"fmt"
	"io"

	"kisumu/pkg/parser"
)

const PROMPT = "kisumu> "

// Start is a read-eval-print loop (REPL) for the Kisumu programming language.
// It continuously reads input from the provided reader, tokenizes it,
// and writes the token type and literal to the provided writer.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	writer := bufio.NewWriter(out)

	for {
		fmt.Fprintln(writer, PROMPT)
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
		lexer := parser.Tokenize(line)

		for {
			tok := lexer.GetNextToken()
			if tok.Type == parser.EOF {
				break
			}
			if tok.Type == parser.ILLEGAL {
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
	lexer  *parser.Lexer
	// ksm    *lexer
	// parser *parser
	// env    *Environment
	// runner *runner
	input  *bufio.Reader
	output *bufio.Writer
}

func NewREPL(prompt string, lexer *parser.Lexer, input *bufio.Reader, output *bufio.Writer) *REPL {
	return &REPL{
		prompt: prompt,
		lexer:  lexer,
		// ksm:    ksm,
		// parser: parser,
		// env:    env,
		// runner: runner,
		input:  input,
		output: output,
	}
}
