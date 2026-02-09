package token

// TokenType represents the type of a token
type TokenType string

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Token types
const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + literals
	IDENT  TokenType = "IDENT"  // variable names
	INT    TokenType = "INT"    // 12345
	STRING TokenType = "STRING" // "hello"

	// Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERISK TokenType = "*"
	SLASH    TokenType = "/"
	MODULO   TokenType = "%"

	BANG TokenType = "!"

	LT  TokenType = "<"
	GT  TokenType = ">"
	LTE TokenType = "<="
	GTE TokenType = ">="
	EQ  TokenType = "=="
	NEQ TokenType = "!="

	AND TokenType = "&&"
	OR  TokenType = "||"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	NEWLINE   TokenType = "NEWLINE"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// Keywords (Sasak language)
	GAWE   TokenType = "GAWE"   // var/let
	TETEP  TokenType = "TETEP"  // const
	YEN    TokenType = "YEN"    // if
	NENG   TokenType = "NENG"   // else
	SALAMA TokenType = "SALAMA" // while
	KANGGO TokenType = "KANGGO" // for
	PUNGSI TokenType = "PUNGSI" // function
	BALIK  TokenType = "BALIK"  // return

	// Boolean and null literals
	BENER  TokenType = "BENER"  // true
	SALAH  TokenType = "SALAH"  // false
	KOSONG TokenType = "KOSONG" // null
)

// keywords maps Sasak keywords to their token types
// Note: tulis, tanya are NOT keywords - they are built-in functions
var keywords = map[string]TokenType{
	"gawe":   GAWE,
	"tetep":  TETEP,
	"yen":    YEN,
	"neng":   NENG,
	"salama": SALAMA,
	"kanggo": KANGGO,
	"pungsi": PUNGSI,
	"balik":  BALIK,
	"bener":  BENER,
	"salah":  SALAH,
	"kosong": KOSONG,
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
