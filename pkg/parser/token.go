package parser

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct { // position and readPosition are used to access characters in input(as index)
	input        string
	position     int  // index of the starting position of the current token(the previous character)
	readPosition int  // index of the current character
	currentChar  rune // current character under examination
}

const (
	ILLEGAL TokenType = "IILEGAL" // Invalid token/unknown character
	EOF     TokenType = "EOF"     // End of file

	// Identifiers and literals
	INT        TokenType = "INT"        // 1323145567890
	STRING     TokenType = "STRING"     // concatenate, slice, and get
	IDENTIFIER TokenType = "IDENTIFIER" // variable name, function name, or struct name

	// Operators and delimiters
	OPEN_BRACKET      TokenType = "OPEN_BRACKET"      // [
	CLOSE_BRACKET     TokenType = "CLOSE_BRACKET"     // ]
	OPEN_CURLY        TokenType = "OPEN_CURLY"        // {
	CLOSE_CURLY       TokenType = "CLOSE_CURLY"       // }
	OPEN_PARENTHESES  TokenType = "OPEN_PARENTHESES"  // (
	CLOSE_PARENTHESES TokenType = "CLOSE_PARENTHESES" // )

	ASSIGNMENT TokenType = "ASSIGNMENT" // =
	EQUALS     TokenType = "EQUALS"     // ==
	NOT        TokenType = "NOT"
	NOT_EQUALS TokenType = "NOT_EQUALS" // !=

	LESS           TokenType = "LESS"           // <
	LESS_EQUAL     TokenType = "LESS_EQUAL"     // <=
	GREATER        TokenType = "GREATER"        // >
	GREATER_EQUALS TokenType = "GREATER_EQUALS" // >=

	OR  TokenType = "OR"  // ||
	AND TokenType = "AND" // &&

	NULL  TokenType = "NULL"  // null
	TRUE  TokenType = "TRUE"  // true
	FALSE TokenType = "FALSE" // false

	DOT        TokenType = "DOT"        //.
	DOT_DOT    TokenType = "DOT_DOT"    //..
	SEMI_COLON TokenType = "SEMI_COLON" // ;
	COLON      TokenType = "COLON"      // :
	QUESTION   TokenType = "QUESTION"   //?
	COMMA      TokenType = "COMMA"      //,
	WHITESPACE TokenType = "WHITESPACE" // Whitespace

	PLUS_PLUS    TokenType = "PLUS_PLUS"    // ++
	MINUS_MINUS  TokenType = "MINUS_MINUS"  // --
	PLUS_EQUALS  TokenType = "PLUS_EQUALS"  // +=
	MINUS_EQUALS TokenType = "MINUS_EQUALS" // -=
	SLASH_EQUALS TokenType = "SLASH_EQUALS" // /=
	STAR_EQUALS  TokenType = "STAR_EQUALS"  // *=

	PLUS    TokenType = "PLUS"    // +
	DASH    TokenType = "DASH"    // -
	SLASH   TokenType = "SLASH"   // /
	STAR    TokenType = "STAR"    // *
	PERCENT TokenType = "PERCENT" // %

	/* ====== RESERVED KEYWORDS ======= */
	LET     TokenType = "LET"     // let
	CONST   TokenType = "CONST"   // const
	CLASS   TokenType = "CLASS"   // class
	NEW     TokenType = "NEW"     // new
	IMPORT  TokenType = "IMPORT"  // import
	FROM    TokenType = "FROM"    // from
	FN      TokenType = "FN"      // fn
	IF      TokenType = "IF"      // if
	ELSE    TokenType = "ELSE"    // else
	FOREACH TokenType = "FOREACH" // foreach
	WHILE   TokenType = "WHILE"   // while
	FOR     TokenType = "FOR"     // for
	EXPORT  TokenType = "EXPORT"  // export
	TYPEOF  TokenType = "TYPEOF"  // typeof
	IN      TokenType = "IN"      // in
)

func (token Token) isAmongDefined(expectedTokens ...TokenType) bool {
	for _, expected := range expectedTokens {
		if expected == token.Type {
			return true
		}
	}
	return false
}

// HELPER METHODS FOR DEBUGGING TOKENS
/*
 Inside the Debug method, it checks the Type of the token. If the token kind is IDENTIFIER, NUMBER, or STRING, it prints the token kind followed by its value in parentheses.
 For example, if the token is an identifier with the value "example", it would print "identifier(example)".
*/
func (token Token) Debug() {
	if token.isAmongDefined(IDENTIFIER, INT, STRING) {
		fmt.Printf("%s (%s)\n", TokenType(token.Type), token.Literal)
	} else {
		fmt.Printf("%s ()\n", TokenType(token.Type))
	}
}

// Function used to create a lexer
func NewToken(Type TokenType, value string) Token {
	return Token{
		Type, value,
	}
}

func Tokenize(input string) *Lexer {
	tok := &Lexer{input: input}
	tok.getNextChar()
	return tok
}

// getNextChar advances the lexer to the next character in the input string.
// It updates the current character (currentChar), the current position (position),
// and the read position (readPosition). If the read position is at or beyond the end of the input string,
// the current character is set to 0.
func (tokens *Lexer) getNextChar() {
	if tokens.readPosition >= len(tokens.input) {
		tokens.currentChar = 0 // ASCII code -> NULL
	} else {
		tokens.currentChar = rune(tokens.input[tokens.readPosition])
	}
	tokens.position = tokens.readPosition
	tokens.readPosition++
}
