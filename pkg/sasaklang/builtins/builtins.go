package builtins

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/object"
)

// builtinTedem sleeps for n milliseconds
func builtinTedem(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("tedem() butuh 1 argumen (ms), dapat %d", len(args))}
	}

	arg, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "argumen tedem() harus angka"}
	}

	time.Sleep(time.Duration(arg.Value) * time.Millisecond)
	return &object.Null{}
}

// builtinAcak returns a random number between 0 and n-1
func builtinAcak(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("acak() butuh 1 argumen (max), dapat %d", len(args))}
	}

	arg, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "argumen acak() harus angka"}
	}

	if arg.Value <= 0 {
		return &object.Error{Message: "argumen acak() harus lebih besar dari 0"}
	}

	return &object.Integer{Value: int64(rand.Int63n(arg.Value))}
}

// Builtins contains all builtin functions
var Builtins = map[string]*object.Builtin{
	"cetak":  {Fn: builtinCetak},
	"isik":   {Fn: builtinIsik},
	"belong": {Fn: builtinBelong},
	"jenis":  {Fn: builtinJenis},
	"waktu":  {Fn: builtinWaktu},
	"sorong": {Fn: builtinSorong},
	"bait":   {Fn: builtinBait},   // get -> bait
	"ngatur": {Fn: builtinNgatur}, // set -> ngatur
	"tedem":  {Fn: builtinTedem},
	"acak":   {Fn: builtinAcak},
}

// builtinCetak prints arguments separated by space with newline
func builtinCetak(args ...object.Object) object.Object {
	strs := make([]string, len(args))
	for i, arg := range args {
		strs[i] = arg.Inspect()
	}
	fmt.Println(strings.Join(strs, " "))
	return &object.Null{}
}

// builtinIsik reads input from stdin
func builtinIsik(args ...object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "isik() butuh maksimal 1 argumen"}
	}

	if len(args) == 1 {
		fmt.Print(args[0].Inspect())
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return &object.Error{Message: "gagal membaca input"}
	}

	return &object.String{Value: strings.TrimSuffix(input, "\n")}
}

// builtinBelong returns the length of string or array
func builtinBelong(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("belong() butuh 1 argumen, dapat %d", len(args))}
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return &object.Error{Message: fmt.Sprintf("belong() tidak mendukung tipe %s", arg.Type())}
	}
}

// builtinJenis returns the type name of an object
func builtinJenis(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("jenis() butuh 1 argumen, dapat %d", len(args))}
	}

	var typeName string
	switch args[0].(type) {
	case *object.Integer:
		typeName = "angka"
	case *object.String:
		typeName = "teks"
	case *object.Boolean:
		typeName = "boolean"
	case *object.Null:
		typeName = "ndarak"
	case *object.Array:
		typeName = "daftar"
	case *object.Map:
		typeName = "peta"
	case *object.Function:
		typeName = "fungsi"
	case *object.Builtin:
		typeName = "fungsi_bawaan"
	default:
		typeName = "tidak_dikenal"
	}

	return &object.String{Value: typeName}
}

// builtinWaktu returns the current unix timestamp
func builtinWaktu(args ...object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "waktu() tidak butuh argumen"}
	}
	return &object.Integer{Value: time.Now().Unix()}
}

// builtinSorong appends an element to an array
func builtinSorong(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("sorong() butuh 2 argumen, dapat %d", len(args))}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "argumen pertama sorong() harus daftar"}
	}

	newElements := make([]object.Object, len(arr.Elements)+1)
	copy(newElements, arr.Elements)
	newElements[len(arr.Elements)] = args[1]

	return &object.Array{Elements: newElements}
}

// builtinBait returns an element from array or map (get)
func builtinBait(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("bait() butuh 2 argumen (koleksi, indeks/kunci), dapat %d", len(args))}
	}

	switch container := args[0].(type) {
	case *object.Array:
		idx, ok := args[1].(*object.Integer)
		if !ok {
			return &object.Error{Message: "indeks harus angka"}
		}
		if idx.Value < 0 || idx.Value >= int64(len(container.Elements)) {
			return &object.Null{}
		}
		return container.Elements[idx.Value]

	case *object.Map:
		key, ok := args[1].(object.Hashable)
		if !ok {
			return &object.Error{Message: "kunci peta tidak valid"}
		}
		pair, ok := container.Pairs[key.HashKey()]
		if !ok {
			return &object.Null{}
		}
		return pair.Value

	default:
		return &object.Error{Message: "argumen pertama harus daftar atau peta"}
	}
}

// builtinNgatur sets an element in array or map (set) - returns the modified collection
// Note: SasakLang objects are immutable by default in implementation (except map/array internal pointers),
// but builtins usually return new objects or modifying them?
// The existing implementation of `evalIndexExpression` supports reading.
// Writing via index `arr[0] = 1` involves `evalAssignmentExpression` but that likely needs AST support for index assignment which parser might not handle as "assignment" (only IDENT).
// Parser `parseAssignmentExpression` takes an identifier.
// So `ngatur` is essential for modifying arrays/maps if syntax doesn't support `arr[i] = val`.
func builtinNgatur(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: fmt.Sprintf("ngatur() butuh 3 argumen (koleksi, indeks/kunci, nilai), dapat %d", len(args))}
	}

	switch container := args[0].(type) {
	case *object.Array:
		idx, ok := args[1].(*object.Integer)
		if !ok {
			return &object.Error{Message: "indeks harus angka"}
		}

		// For simplicity, we modify correctly. If index out of bounds, maybe grow or error?
		if idx.Value < 0 || idx.Value >= int64(len(container.Elements)) {
			return &object.Error{Message: "indeks di luar batas"}
		}

		container.Elements[idx.Value] = args[2] // Mutation
		return container

	case *object.Map:
		key, ok := args[1].(object.Hashable)
		if !ok {
			return &object.Error{Message: "kunci peta tidak valid"}
		}
		container.Pairs[key.HashKey()] = object.MapPair{Key: args[1], Value: args[2]} // Mutation
		return container

	default:
		return &object.Error{Message: "argumen pertama harus daftar atau peta"}
	}
}
