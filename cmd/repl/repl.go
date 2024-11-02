package repl

import (
	"bufio"
	"fmt"
	"io"

	"kisumu/pkg/interpreter"
	"kisumu/pkg/lexer"
	"kisumu/pkg/object"
	"kisumu/pkg/parser"
)

const PROMPT = "kisumu $"

// Start is a read-eval-print loop (REPL) for the Kisumu programming language.
// It continuously reads input from the provided reader, tokenizes it,
// and writes the token type and literal to the provided writer.
func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)

		if !scanner.Scan() {
			return // Exit loop on EOF or error
		}

		line := scanner.Text()
		Lexer := lexer.Tokenize(line)
		parser := parser.NewParser(Lexer)

		program := parser.ParseProgram()
		if len(parser.Errors()) != 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		evaluated := interpreter.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
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
