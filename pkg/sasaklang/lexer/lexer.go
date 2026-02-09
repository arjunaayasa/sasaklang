package lexer

import (
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/token"
)

// Lexer tokenizes source code
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readChar()
	return l
}

// readChar reads the next character and advances position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++

	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

// peekChar returns the next character without advancing
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Save position for token
	line := l.line
	column := l.column

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "==", Line: line, Column: column}
		} else {
			tok = newToken(token.ASSIGN, l.ch, line, column)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch, line, column)
	case '-':
		tok = newToken(token.MINUS, l.ch, line, column)
	case '*':
		tok = newToken(token.ASTERISK, l.ch, line, column)
	case '/':
		tok = newToken(token.SLASH, l.ch, line, column)
	case '%':
		tok = newToken(token.MODULO, l.ch, line, column)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NEQ, Literal: "!=", Line: line, Column: column}
		} else {
			tok = newToken(token.BANG, l.ch, line, column)
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: "<=", Line: line, Column: column}
		} else {
			tok = newToken(token.LT, l.ch, line, column)
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: ">=", Line: line, Column: column}
		} else {
			tok = newToken(token.GT, l.ch, line, column)
		}
	case '&':
		if l.peekChar() == '&' {
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: "&&", Line: line, Column: column}
		} else {
			tok = newToken(token.ILLEGAL, l.ch, line, column)
		}
	case '|':
		if l.peekChar() == '|' {
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: "||", Line: line, Column: column}
		} else {
			tok = newToken(token.ILLEGAL, l.ch, line, column)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch, line, column)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, line, column)
	case ':':
		tok = newToken(token.COLON, l.ch, line, column)
	case '(':
		tok = newToken(token.LPAREN, l.ch, line, column)
	case ')':
		tok = newToken(token.RPAREN, l.ch, line, column)
	case '{':
		tok = newToken(token.LBRACE, l.ch, line, column)
	case '}':
		tok = newToken(token.RBRACE, l.ch, line, column)
	case '[':
		tok = newToken(token.LBRACKET, l.ch, line, column)
	case ']':
		tok = newToken(token.RBRACKET, l.ch, line, column)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Line = line
		tok.Column = column
		return tok
	case '#':
		l.skipComment()
		return l.NextToken()
	case '\n':
		tok = newToken(token.NEWLINE, l.ch, line, column)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = line
		tok.Column = column
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = line
			tok.Column = column
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Line = line
			tok.Column = column
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, line, column)
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace skips spaces and tabs (but not newlines)
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips a single-line comment
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// readIdentifier reads an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readString reads a string literal
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
		// Handle escape sequences
		if l.ch == '\\' && l.peekChar() != 0 {
			l.readChar()
		}
	}
	str := l.input[position:l.position]
	l.readChar() // consume closing quote
	return str
}

// newToken creates a new token
func newToken(tokenType token.TokenType, ch byte, line, column int) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line, Column: column}
}

// isLetter checks if a character is a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if a character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
