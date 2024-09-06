package repl

import (
	"bufio"
	"fmt"
	"io"

	"kisumu/pkg/parser"
)

const PROMPT = "kisumu> "

// Start is a read-eval-print loop (REPL) for the Kisumu programming language.
// It continuously reads input from the provided reader, parses it into an abstract syntax tree (AST),
// evaluates the AST using the provided environment, and writes the result to the provided writer.
//
// Parameters:
// - in: An io.Reader from which to read input.
// - out: An io.Writer to which to write output.
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
		tree := parser.Tokenize(line)

		for tok := tree.GetNextToken(); tok.Type != parser.EOF; tok = tree.GetNextToken() {
			fmt.Fprintf(writer, "Token: %s (%s)\n", tok.Type, tok.Literal)
		}

		writer.Flush()  // write results to the output
	}
}

type REPL struct {
	prompt string
	ksm    *parser.Lexer
	parser *parser.Lexer
	env    *parser.Lexer
	runner *parser.Lexer
	input  *bufio.Reader
	output *bufio.Writer
}

func NewREPL(prompt string, ksm, parser, env, runner *parser.Lexer, input *bufio.Reader, output *bufio.Writer) *REPL {
	return &REPL{
		prompt: prompt,
		ksm:    ksm,
		parser: parser,
		env:    env,
		runner: runner,
		input:  input,
		output: output,
	}
}

func (r *REPL) Start() {
	for {
		fmt.Fprintf(r.output, r.prompt)
		r.output.Flush() // Ensure prompt is printed before reading input
		scanned, err := r.input.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break // Exit loop on end of input
			}
			fmt.Fprintf(r.output, "Error reading input: %v\n", err)
			continue
		}

		// TODO: Add parsing, evaluating, and writing results here
		fmt.Fprintf(r.output, "You typed: %s", scanned)
		r.output.Flush() // Ensure output is written
	}
}
