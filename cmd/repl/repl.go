package repl

import (
	"bufio"
)

type REPL struct {
	prompt string
	ksm    *lexer
	parser *parser
	env    *Environment
	runner *runner
	input  *bufio.Reader
	output *bufio.Writer
}

func NewREPL(prompt string, ksm, parser, env, runner *lexer, input, output *bufio.Reader) *REPL {
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
