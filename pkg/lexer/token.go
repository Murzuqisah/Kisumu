package lexer

import (
	"fmt"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct { // position and readPosition are used to access characters in input(as index)
	input        string
	position     int  // index of the starting position of the current token(the previous character)
	readPosition int  // index of the current character
	currentChar  byte // current character under examination
}

const (
	ILLEGAL = "ILLEGAL" // Invalid token/unknown character
	EOF     = "EOF"     // End of file

	// Identifiers and literals
	KEYWORD     = "KEYWORD"     // break, continue, else, for, if, return, struct, var
	RETURN_TYPE = "RETURN_TYPE" // int, string, etc.
	STRUCT_TYPE = "STRUCT_TYPE" // struct { field1 type; field2 type; }
	VAR         = "VAR"         // var     // const
	TYPE        = "TYPE"        // int, string, etc.
	BOOLEAN     = "BOOLEAN"     // true, false
	FLOAT       = "FLOAT"       // 123.456
	IMAGINARY   = "IMAGINARY"   // 123.456i
	RUNE        = "RUNE"        // 'a'
	INT         = "INT"         // 1323145567890
	STRING      = "STRING"      // concatenate, slice, and get
	IDENTIFIER  = "IDENTIFIER"  // variable name, function name, or struct name

	// Operators and delimiters
	OPEN_BRACKET      = "OPEN_BRACKET"      // [
	CLOSE_BRACKET     = "CLOSE_BRACKET"     // ]
	OPEN_CURLY        = "OPEN_CURLY"        // {
	CLOSE_CURLY       = "CLOSE_CURLY"       // }
	OPEN_PARENTHESES  = "OPEN_PARENTHESES"  // (
	CLOSE_PARENTHESES = "CLOSE_PARENTHESES" // )

	ASSIGNMENT = "ASSIGNMENT" // =
	EQUALS     = "EQUALS"     // ==
	NOT        = "NOT"
	NOT_EQUALS = "NOT_EQUALS" // !=

	LESS           = "LESS"           // <
	LESS_EQUAL     = "LESS_EQUAL"     // <=
	GREATER        = "GREATER"        // >
	GREATER_EQUALS = "GREATER_EQUALS" // >=

	OR    = "OR"    // ||
	AND   = "AND"   // &&
	NULL  = "NULL"  // null
	TRUE  = "TRUE"  // true
	FALSE = "FALSE" // false

	BANG       = "BANG"
	DOT        = "DOT"        //.
	DOT_DOT    = "DOT_DOT"    //..
	SEMI_COLON = "SEMI_COLON" // ;
	COLON      = "COLON"      // :
	QUESTION   = "QUESTION"   //?
	COMMA      = "COMMA"      //,
	WHITESPACE = "WHITESPACE" // Whitespace

	PLUS_PLUS    = "PLUS_PLUS"    // ++
	MINUS_MINUS  = "MINUS_MINUS"  // --
	PLUS_EQUALS  = "PLUS_EQUALS"  // +=
	MINUS_EQUALS = "MINUS_EQUALS" // -=
	SLASH_EQUALS = "SLASH_EQUALS" // /=
	STAR_EQUALS  = "STAR_EQUALS"  // *=

	PLUS     = "PLUS"     // +
	DASH     = "DASH"     // -
	SLASH    = "SLASH"    // /
	ASTERISK = "ASTERISK" // *
	PERCENT  = "PERCENT"  // %

	/* ====== RESERVED KEYWORDS ======= */
	LET      = "LET"      // let
	CONST    = "CONST"    // const
	CLASS    = "CLASS"    // class
	NEW      = "NEW"      // new
	IMPORT   = "IMPORT"   // import
	FROM     = "FROM"     // from
	FN       = "FUNCTION" // fn
	IF       = "IF"       // if
	ELSE     = "ELSE"     // else
	FOREACH  = "FOREACH"  // foreach
	WHILE    = "WHILE"    // while
	FOR      = "FOR"      // for
	EXPORT   = "EXPORT"   // export
	TYPEOF   = "TYPEOF"   // typeof
	IN       = "IN"       // in
	RETURN   = "RETURN"   // return
	BREAK    = "BREAK"    // break
	CONTINUE = "CONTINUE" // continue

)

var KEYWORDS = map[string]TokenType{
	"function":  FN,
	"let":       LET,
	"const":     CONST,
	"class":     CLASS,
	"new":       NEW,
	"import":    IMPORT,
	"from":      FROM,
	"fn":        FN,
	"if":        IF,
	"else":      ELSE,
	"foreach":   FOREACH,
	"while":     WHILE,
	"for":       FOR,
	"export":    EXPORT,
	"typeof":    TYPEOF,
	"in":        IN,
	"return":    RETURN,
	"break":     BREAK,
	"continue":  CONTINUE,
	"null":      NULL,
	"true":      TRUE,
	"false":     FALSE,
	"boolean":   BOOLEAN,
	"float":     FLOAT,
	"imaginary": IMAGINARY,
	"rune":      RUNE,
	"int":       INT,
	"string":    STRING,
	"struct":    STRUCT_TYPE,
	"var":       VAR,
	"type":      TYPE,
	"or":        OR,
	"and":       AND,
}

func (token Token) isAmongDefined(expectedTokens ...TokenType) bool {
	for _, expected := range expectedTokens {
		if expected == token.Type {
			return true
		}
	}
	return false
}

func LookupIdentifier(keyword string) TokenType {
	if keywordToken, ok := KEYWORDS[keyword]; ok {
		return keywordToken
	}
	return IDENTIFIER
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

func IsLetter(ch byte) bool {

	return 'a' <= ch && ch >= 'z' || 'A' <= ch && ch >= 'Z' || ch == '_'
}

func IsFloat(s string) bool {
	for _, char := range s {
		if char == '.' {
			return true
		}
		if !IsDigit(byte(char)) {
			return false
		}
	}
	return false
}

// readIdentifier extracts a sequence of characters from the input that form an identifier.
// It continues reading characters until it encounters a non-identifier character (e.g., whitespace or punctuation).
//
// Parameters:
// - l: A pointer to the Lexer struct. This function is a method of the Lexer struct, so it has access to its fields and methods.
//
// Returns:
// - A string representing the extracted identifier.
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for IsLetter(byte(l.currentChar)) || l.currentChar == '_' {
		l.getChar()
	}
	return l.input[startPosition:l.position]
}

// Function used to create a lexer
func NewToken(Type TokenType, value string) Token {
	return Token{
		Type, value,
	}
}

// Tokenize initializes and returns a new Lexer instance with the given input string.
// The Lexer is responsible for breaking down the input string into a sequence of tokens.
//
// Parameters:
// - input: A string containing the source code to be tokenized.
//
// Returns:
// - A pointer to a new Lexer instance, initialized with the provided input string.
func Tokenize(input string) *Lexer {
	tok := &Lexer{input: input} // Create a new Lexer instance with the given input string
	tok.getChar()               // Initialize the current character and position of the lexer
	return tok                  // Return the initialized Lexer instance
}

// getNextChar advances the lexer to the next character in the input string.
// It updates the current character (currentChar), the current position (position),
// and the read position (readPosition). If the read position is at or beyond the end of the input string,
// the current character is set to 0.
func (tokens *Lexer) getChar() { //readChar() advances the lexer to the next character in the input string
	if tokens.readPosition >= len(tokens.input) {
		tokens.currentChar = 0 // ASCII code -> NULL
	} else {
		tokens.currentChar = tokens.input[tokens.readPosition]
	}
	tokens.position = tokens.readPosition
	tokens.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 // End of input
	}
	return l.input[l.readPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.getChar()
	}
}

func (l *Lexer) readNumber() string {
	var number string

	for IsDigit(byte(l.currentChar)) {
		number += string(l.currentChar)
		l.getChar() // read & fetch the character
	}

	if l.currentChar == '.' {
		number += string(l.currentChar)
		l.getChar() // read & fetch the character

		for IsDigit(byte(l.currentChar)) {
			number += string(l.currentChar)
		}
	}
	return number
}

func IsDigit(char byte) bool {
	return '0' <= char && char <= '9'

}

func (l *Lexer) GetNextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.currentChar {
	case '[':
		tok = newToken(OPEN_BRACKET, string(l.currentChar))
	case ']':
		tok = newToken(CLOSE_BRACKET, string(l.currentChar))
	case '{':
		tok = newToken(OPEN_CURLY, string(l.currentChar))
	case '}':
		tok = newToken(CLOSE_CURLY, string(l.currentChar))
	case '(':
		tok = newToken(OPEN_PARENTHESES, string(l.currentChar))
	case ')':
		tok = newToken(CLOSE_PARENTHESES, string(l.currentChar))
	case '=':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.getChar()
			tok = newToken(EQUALS, string(ch)+string(l.currentChar))
		} else {
			tok = newToken(ASSIGNMENT, string(l.currentChar))
		}
	case '+':
		if l.peekChar() == '+' {
			l.getChar()
			tok = newToken(PLUS_PLUS, "++")
		} else if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(PLUS_EQUALS, "+=")
		} else {
			tok = newToken(PLUS, string(l.currentChar))
		}
	case '-':
		if l.peekChar() == '-' {
			l.getChar()
			tok = newToken(MINUS_MINUS, "--")
		} else if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(MINUS_EQUALS, "-=")
		} else {
			tok = newToken(DASH, string(l.currentChar))
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.currentChar
			l.getChar()
			tok = newToken(NOT_EQUALS, string(ch)+string(l.currentChar))
		} else {
			tok = newToken(BANG, string(l.currentChar))
		}
	case '*':
		if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(STAR_EQUALS, "*=")
		} else {
			tok = newToken(ASTERISK, string(l.currentChar))
		}
	case '/':
		if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(SLASH_EQUALS, "/=")
		} else {
			tok = newToken(SLASH, string(l.currentChar))
		}
	case '<':
		if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(LESS_EQUAL, "<=")
		} else {
			tok = newToken(LESS, string(l.currentChar))
		}
	case '>':
		if l.peekChar() == '=' {
			l.getChar()
			tok = newToken(GREATER_EQUALS, ">=")
		} else {
			tok = newToken(GREATER, string(l.currentChar))
		}
	case '|':
		if l.peekChar() == '|' {
			l.getChar()
			tok = newToken(OR, "||")
		}
	case '&':
		if l.peekChar() == '&' {
			l.getChar()
			tok = newToken(AND, "&&")
		}
	case '.':
		if l.peekChar() == '.' {
			l.getChar()
			if l.peekChar() == '.' {
				l.getChar()
				tok = newToken(DOT_DOT, "...")
			} else {
				tok = newToken(DOT, string(l.currentChar))
			}
		} else {
			tok = newToken(DOT, string(l.currentChar))
		}
	case ';':
		tok = newToken(SEMI_COLON, string(l.currentChar))
	case ':':
		tok = newToken(COLON, string(l.currentChar))
	case '?':
		tok = newToken(QUESTION, string(l.currentChar))
	case ',':
		tok = newToken(COMMA, string(l.currentChar))
	case 0:
		tok = newToken(EOF, "")
	default:
		if IsLetter(l.currentChar) || l.currentChar == '_' {
			ident := l.readIdentifier()
			tok = NewToken(LookupIdentifier(ident), ident) // Use LookupIdentifier here
			return tok                                     // Return here to prevent getting the next character too early
		} else if IsDigit(l.currentChar) {
			tok = newToken(INT, l.readNumber())
			return tok // Return here to prevent getting the next character too early
		} else {
			tok = newToken(ILLEGAL, string(l.currentChar))
		}
	}

	l.getChar() // Move to the next character only if we haven't already returned a token
	return tok
}

func newToken(tokenType TokenType, currentChar string) Token {
	return Token{
		Type:    tokenType,
		Literal: currentChar,
	}
}
