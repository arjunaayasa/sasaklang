package lexer

import (
	"testing"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/token"
)

func TestNextToken(t *testing.T) {
	input := `gawe x = 5
tetep PI = 3
tulis("hello")
tanya("nama: ")

yen (x > 3) {
    balik bener
} neng {
    balik salah
}

salama (x > 0) {
    x = x - 1
}

kanggo (gawe i = 0; i < 10; i = i + 1) {
    tulis(i)
}

pungsi tambah(a, b) {
    balik a + b
}

# ini komentar
gawe arr = [1, 2, 3]
gawe obj = {"nama": "test"}

kosong
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.GAWE, "gawe"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NEWLINE, "\n"},
		{token.TETEP, "tetep"},
		{token.IDENT, "PI"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "tulis"},
		{token.LPAREN, "("},
		{token.STRING, "hello"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "tanya"},
		{token.LPAREN, "("},
		{token.STRING, "nama: "},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.YEN, "yen"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.BALIK, "balik"},
		{token.BENER, "bener"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NENG, "neng"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.BALIK, "balik"},
		{token.SALAH, "salah"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.SALAMA, "salama"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.IDENT, "x"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.KANGGO, "kanggo"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q (literal=%q)",
				i, tt.expectedType, tok.Type, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestOperators(t *testing.T) {
	input := `+ - * / % = == != < > <= >= && || !`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, "/"},
		{token.MODULO, "%"},
		{token.ASSIGN, "="},
		{token.EQ, "=="},
		{token.NEQ, "!="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.AND, "&&"},
		{token.OR, "||"},
		{token.BANG, "!"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLineColumn(t *testing.T) {
	input := `gawe x = 5
gawe y = 10`

	l := New(input)

	// gawe
	tok := l.NextToken()
	if tok.Line != 1 {
		t.Errorf("expected line 1, got %d", tok.Line)
	}

	// skip to second line
	for tok.Type != token.NEWLINE {
		tok = l.NextToken()
	}

	// gawe on second line
	tok = l.NextToken()
	if tok.Line != 2 {
		t.Errorf("expected line 2, got %d", tok.Line)
	}
}

func TestComments(t *testing.T) {
	input := `gawe x = 5 # ini komentar
gawe y = 10`

	l := New(input)

	expected := []token.TokenType{
		token.GAWE,
		token.IDENT,
		token.ASSIGN,
		token.INT,
		token.NEWLINE,
		token.GAWE,
		token.IDENT,
		token.ASSIGN,
		token.INT,
		token.EOF,
	}

	for i, exp := range expected {
		tok := l.NextToken()
		if tok.Type != exp {
			t.Fatalf("tests[%d] - expected %q, got %q", i, exp, tok.Type)
		}
	}
}
