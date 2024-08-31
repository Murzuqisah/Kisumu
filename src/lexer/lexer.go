package lexer

import (
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)


type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

func Tokenize(source string) []Token {
	lex := createLexer(source)

	return lex.Tokens
}

// createLexer initializes and returns a new lexer instance for the given source string.
//
// Parameters:
//   - source: A string containing the input text to be tokenized.
//
// Returns:
//   A pointer to a new lexer instance, initialized with:
//   - The starting position set to 0.
//   - The source string.
//   - An empty slice of Tokens.
//   - A predefined set of regex patterns and their corresponding handlers.
func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			/* ======== DEFINE ALL THE PATTERNS =========== */
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
		},
	}
}

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.Tokens = append(lex.Tokens, NewToken(kind, value))
		lex.pos += len(regex.FindString(lex.source[lex.pos:]))
	}
}
