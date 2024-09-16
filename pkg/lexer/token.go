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
	currentChar  rune // current character under examination
}

const (
	ILLEGAL TokenType = "ILLEGAL" // Invalid token/unknown character
	EOF     TokenType = "EOF"     // End of file

	// Identifiers and literals
	KEYWORD     TokenType = "KEYWORD"     // break, continue, else, for, if, return, struct, var
	FUNCTION    TokenType = "FUNCTION"    // function
	RETURN_TYPE TokenType = "RETURN_TYPE" // int, string, etc.
	STRUCT_TYPE TokenType = "STRUCT_TYPE" // struct { field1 type; field2 type; }
	VAR         TokenType = "VAR"         // var     // const
	TYPE        TokenType = "TYPE"        // int, string, etc.
	BOOLEAN     TokenType = "BOOLEAN"     // true, false
	FLOAT       TokenType = "FLOAT"       // 123.456
	IMAGINARY   TokenType = "IMAGINARY"   // 123.456i
	RUNE        TokenType = "RUNE"        // 'a'
	INT         TokenType = "INT"         // 1323145567890
	STRING      TokenType = "STRING"      // concatenate, slice, and get
	IDENTIFIER  TokenType = "IDENTIFIER"  // variable name, function name, or struct name

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

	BANG       TokenType = "BANG"
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

	PLUS     TokenType = "PLUS"     // +
	DASH     TokenType = "DASH"     // -
	SLASH    TokenType = "SLASH"    // /
	ASTERISK TokenType = "ASTERISK" // *
	PERCENT  TokenType = "PERCENT"  // %

	/* ====== RESERVED KEYWORDS ======= */
	LET      TokenType = "LET"      // let
	CONST    TokenType = "CONST"    // const
	CLASS    TokenType = "CLASS"    // class
	NEW      TokenType = "NEW"      // new
	IMPORT   TokenType = "IMPORT"   // import
	FROM     TokenType = "FROM"     // from
	FN       TokenType = "FN"       // fn
	IF       TokenType = "IF"       // if
	ELSE     TokenType = "ELSE"     // else
	FOREACH  TokenType = "FOREACH"  // foreach
	WHILE    TokenType = "WHILE"    // while
	FOR      TokenType = "FOR"      // for
	EXPORT   TokenType = "EXPORT"   // export
	TYPEOF   TokenType = "TYPEOF"   // typeof
	IN       TokenType = "IN"       // in
	RETURN   TokenType = "RETURN"   // return
	BREAK    TokenType = "BREAK"    // break
	CONTINUE TokenType = "CONTINUE" // continue

)

var KEYWORDS = map[string]TokenType{
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

// readIdentifier extracts a sequence of characters from the input that form an identifier.
// It continues reading characters until it encounters a non-identifier character (e.g., whitespace or punctuation).
//
// Parameters:
// - l: A pointer to the Lexer struct. This function is a method of the Lexer struct, so it has access to its fields and methods.
//
// Returns:
// - A string representing the extracted identifier.

func IsLetter(s string) bool {
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			return true
		}
	}
	return false
}

func IsFloat(s string) bool {
	for _, char := range s {
		if char == '.' {
			return true
		}
		if !IsDigit(string(char)) {
			return false
		}
	}
	return false
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for IsLetter(string(l.currentChar)) || l.currentChar == '_' {
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
func (tokens *Lexer) getChar() {
	if tokens.readPosition >= len(tokens.input) {
		tokens.currentChar = 0 // ASCII code -> NULL
	} else {
		tokens.currentChar = rune(tokens.input[tokens.readPosition])
	}
	tokens.position = tokens.readPosition
	tokens.readPosition++
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0 // End of input
	}
	return rune(l.input[l.readPosition])
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == '\\' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.getChar()
	}
}

func (l *Lexer) readNumber() string {
	var number string

	for IsDigit(string(l.currentChar)) {
		number += string(l.currentChar)
		l.getChar() // read & fetch the character
	}

	if l.currentChar == '.' {
		number += string(l.currentChar)
		l.getChar() // read & fetch the character

		for IsDigit(string(l.currentChar)) {
			number += string(l.currentChar)
		}
	}
	return number
}

func IsDigit(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
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
			l.getChar()
			tok = newToken(EQUALS, "==")
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
			l.getChar()
			tok = newToken(NOT_EQUALS, "!=")
		} else {
			tok = newToken(NOT, string(l.currentChar))
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
		if IsLetter(string(l.currentChar)) || l.currentChar == '_' {
			ident := l.readIdentifier()
			// tok = LookupKeyword(string(ident))
			tok = newToken(IDENTIFIER, ident)
		} else if IsDigit(string(l.currentChar)) {
			tok = newToken(INT, l.readNumber())
		} else if l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
			tok = newToken(WHITESPACE, string(l.currentChar))
		} else {
			tok = newToken(ILLEGAL, string(l.currentChar))
		}
	}

	l.getChar()
	return tok
}

func newToken(tokenType TokenType, currentChar string) Token {
	return Token{
		Type:    tokenType,
		Literal: currentChar,
	}
}
