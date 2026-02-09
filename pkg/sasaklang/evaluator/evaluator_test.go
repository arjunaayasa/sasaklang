package evaluator

import (
	"testing"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/lexer"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/object"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"10 % 3", 1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"kenak", true},
		{"salak", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 <= 1", true},
		{"1 >= 1", true},
		{"1 <= 2", true},
		{"2 >= 1", true},
		{"kenak == kenak", true},
		{"salak == salak", true},
		{"kenak == salak", false},
		{"kenak != salak", true},
		{"(1 < 2) == kenak", true},
		{"(1 < 2) == salak", false},
		{"(1 > 2) == kenak", false},
		{"(1 > 2) == salak", true},
		{"kenak ance kenak", true},
		{"kenak ance salak", false},
		{"salak atau kenak", true},
		{"salak atau salak", false},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"ndek kenak", false},
		{"ndek salak", true},
		{"ndek 5", false},
		{"ndek ndek kenak", true},
		{"ndek ndek salak", false},
		{"ndek ndek 5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"lamun (kenak) { 10 }", 10},
		{"lamun (salak) { 10 }", nil},
		{"lamun (1) { 10 }", 10},
		{"lamun (1 < 2) { 10 }", 10},
		{"lamun (1 > 2) { 10 }", nil},
		{"lamun (1 > 2) { 10 } endah { 20 }", 20},
		{"lamun (1 < 2) { 10 } endah { 20 }", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"gawe a = 5; a;", 5},
		{"gawe a = 5 * 5; a;", 25},
		{"gawe a = 5; gawe b = a; b;", 5},
		{"gawe a = 5; gawe b = a; gawe c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fungsi(x) { x + 2 }"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong parameters. Parameters=%+v",
			fn.Parameters)
	}

	if fn.Parameters[0].String() != "x" {
		t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
	}

	expectedBody := "(x + 2)"
	if fn.Body.String() != expectedBody {
		t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"gawe identity = fungsi(x) { x }; identity(5);", 5},
		{"gawe identity = fungsi(x) { tulakan x }; identity(5);", 5},
		{"gawe double = fungsi(x) { x * 2 }; double(5);", 10},
		{"gawe add = fungsi(x, y) { x + y }; add(5, 5);", 10},
		{"gawe add = fungsi(x, y) { x + y }; add(5 + 5, add(5, 5));", 20},
		{"fungsi(x) { x }(5)", 5},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestNamedFunction(t *testing.T) {
	input := `
fungsi tambah(a, b) {
    tulakan a + b
}
tambah(3, 4)
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 7)
}

func TestRecursiveFunction(t *testing.T) {
	input := `
fungsi factorial(n) {
    lamun (n <= 1) {
        tulakan 1
    } endah {
        tulakan n * factorial(n - 1)
    }
}
factorial(5)
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 120)
}

func TestClosures(t *testing.T) {
	input := `
gawe newAdder = fungsi(x) {
    fungsi(y) { x + y }
}
gawe addTwo = newAdder(2)
addTwo(2)
`
	testIntegerObject(t, testEval(input), 4)
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"[1, 2, 3][0]", 1},
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][2]", 3},
		{"gawe i = 0; [1][i];", 1},
		{"[1, 2, 3][1 + 1]", 3},
		{"gawe myArray = [1, 2, 3]; myArray[2];", 3},
		{"gawe myArray = [1, 2, 3]; myArray[0] + myArray[1] + myArray[2];", 6},
		{"gawe myArray = [1, 2, 3]; gawe i = myArray[0]; myArray[i]", 2},
		{"[1, 2, 3][3]", nil},
		{"[1, 2, 3][-1]", nil},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestConstReassignError(t *testing.T) {
	input := `
tetep PI = 3
PI = 4
`
	evaluated := testEval(input)
	errObj, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected error object, got=%T (%+v)", evaluated, evaluated)
	}

	expectedMsg := "tidak bisa mengubah konstanta 'PI'"
	if errObj.Message != expectedMsg {
		t.Errorf("wrong error message. expected=%q, got=%q",
			expectedMsg, errObj.Message)
	}
}

func TestDivisionByZeroError(t *testing.T) {
	input := "10 / 0"
	evaluated := testEval(input)
	errObj, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected error object, got=%T (%+v)", evaluated, evaluated)
	}

	expectedMsg := "pembagian dengan nol"
	if errObj.Message != expectedMsg {
		t.Errorf("wrong error message. expected=%q, got=%q",
			expectedMsg, errObj.Message)
	}
}

func TestUndefinedVariableError(t *testing.T) {
	input := "x"
	evaluated := testEval(input)
	errObj, ok := evaluated.(*object.Error)
	if !ok {
		t.Fatalf("expected error object, got=%T (%+v)", evaluated, evaluated)
	}

	if errObj.Message != "variabel 'x' belum didefinisikan" {
		t.Errorf("wrong error message. got=%q", errObj.Message)
	}
}

func TestWhileLoop(t *testing.T) {
	input := `
gawe x = 0
gawe sum = 0
selame (x < 5) {
    sum = sum + x
    x = x + 1
}
sum
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 10) // 0+1+2+3+4 = 10
}

func TestForLoop(t *testing.T) {
	input := `
gawe sum = 0
ojok (gawe i = 1; i <= 5; i = i + 1) {
    sum = sum + i
}
sum
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 15) // 1+2+3+4+5 = 15
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`belong("")`, 0},
		{`belong("four")`, 4},
		{`belong("hello world")`, 11},
		{`belong([1, 2, 3])`, 3},
		{`jenis(5)`, "angka"},
		{`jenis("hello")`, "teks"},
		{`jenis(kenak)`, "boolean"},
		{`jenis(ndarak)`, "ndarak"},
		{`jenis([1, 2])`, "daftar"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			str, ok := evaluated.(*object.String)
			if !ok {
				t.Errorf("object is not String. got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if str.Value != expected {
				t.Errorf("wrong string. expected=%q, got=%q", expected, str.Value)
			}
		}
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}
	return true
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
