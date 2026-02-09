package errors

import "fmt"

// ParseError represents a parsing error
type ParseError struct {
	Message string
	Line    int
	Column  int
	Token   string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse error [baris %d, kolom %d] dekat '%s': %s",
		e.Line, e.Column, e.Token, e.Message)
}

// NewParseError creates a new parse error
func NewParseError(msg string, line, col int, token string) *ParseError {
	return &ParseError{
		Message: msg,
		Line:    line,
		Column:  col,
		Token:   token,
	}
}

// RuntimeError types
const (
	ErrUndefinedVariable = "variabel '%s' belum didefinisikan"
	ErrTypeMismatch      = "tipe tidak cocok: %s %s %s"
	ErrConstReassign     = "tidak bisa mengubah konstanta '%s'"
	ErrDivisionByZero    = "pembagian dengan nol"
	ErrNotAFunction      = "'%s' bukan fungsi"
	ErrWrongArgCount     = "jumlah argumen salah: butuh %d, dapat %d"
	ErrIndexOutOfBounds  = "indeks di luar batas: %d"
	ErrNotIndexable      = "tipe %s tidak bisa diakses dengan indeks"
	ErrUnhashableKey     = "tipe %s tidak bisa digunakan sebagai kunci map"
	ErrInvalidOperator   = "operator tidak valid: %s%s"
	ErrInvalidInfix      = "operator tidak valid: %s %s %s"
)

// FormatError formats a runtime error message
func FormatError(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
