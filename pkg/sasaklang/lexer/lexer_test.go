package lexer

import (
	"testing"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/token"
)

func TestNextToken(t *testing.T) {
	input := `gawe x = 5
tetep PI = 3
cetak("hello")
isik("nama: ")

lamun (x > 3) {
    tulakan kenak
} endah {
    tulakan salak
}

selame (x > 0) {
    x = x - 1
}

ojok (gawe i = 0; i < 10; i = i + 1) {
    cetak(i)
}

fungsi tambah(a, b) {
    tulakan a + b
}

# ini komentar
gawe arr = [1, 2, 3]
gawe obj = {"nama": "test"}

ndarak
mentelah
lanjutan
ance
atau
ndek
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
		{token.IDENT, "cetak"},
		{token.LPAREN, "("},
		{token.STRING, "hello"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "isik"},
		{token.LPAREN, "("},
		{token.STRING, "nama: "},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.YEN, "lamun"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.GT, ">"},
		{token.INT, "3"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.BALIK, "tulakan"},
		{token.BENER, "kenak"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NENG, "endah"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.BALIK, "tulakan"},
		{token.SALAH, "salak"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.SALAMA, "selame"},
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
		{token.KANGGO, "ojok"},
		{token.LPAREN, "("},
		{token.GAWE, "gawe"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.ASSIGN, "="},
		{token.IDENT, "i"},
		{token.PLUS, "+"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.IDENT, "cetak"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.PUNGSI, "fungsi"},
		{token.IDENT, "tambah"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.BALIK, "tulakan"},
		{token.IDENT, "a"},
		{token.PLUS, "+"},
		{token.IDENT, "b"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.GAWE, "gawe"},
		{token.IDENT, "arr"},
		{token.ASSIGN, "="},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.COMMA, ","},
		{token.INT, "3"},
		{token.RBRACKET, "]"},
		{token.NEWLINE, "\n"},
		{token.GAWE, "gawe"},
		{token.IDENT, "obj"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.STRING, "nama"},
		{token.COLON, ":"},
		{token.STRING, "test"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.KOSONG, "ndarak"},
		{token.NEWLINE, "\n"},
		{token.TIPUQ, "mentelah"},
		{token.NEWLINE, "\n"},
		{token.LANJUT, "lanjutan"},
		{token.NEWLINE, "\n"},
		{token.AND, "ance"},
		{token.NEWLINE, "\n"},
		{token.OR, "atau"},
		{token.NEWLINE, "\n"},
		{token.BANG, "ndek"},
		{token.NEWLINE, "\n"},
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
	input := `+ - * / % = == != < > <= >= ance || atau ! ndek`

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
		{token.AND, "ance"},
		{token.OR, "||"},
		{token.OR, "atau"},
		{token.BANG, "!"},
		{token.BANG, "ndek"},
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
