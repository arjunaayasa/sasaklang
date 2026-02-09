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

// builtinAntos sleeps for n milliseconds
func builtinAntos(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("antos() butuh 1 argumen (ms), dapat %d", len(args))}
	}

	arg, ok := args[0].(*object.Integer)
	if !ok {
		return &object.Error{Message: "argumen antos() harus angka"}
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
	"tulis":   {Fn: builtinTulis},
	"tanya":   {Fn: builtinTanya},
	"panjang": {Fn: builtinPanjang},
	"jenis":   {Fn: builtinJenis},
	"waktu":   {Fn: builtinWaktu},
	"dorong":  {Fn: builtinDorong},
	"pertama": {Fn: builtinPertama},
	"akhir":   {Fn: builtinAkhir},
	"antos":   {Fn: builtinAntos},
	"acak":    {Fn: builtinAcak},
}

// builtinTulis prints arguments separated by space with newline
func builtinTulis(args ...object.Object) object.Object {
	strs := make([]string, len(args))
	for i, arg := range args {
		strs[i] = arg.Inspect()
	}
	fmt.Println(strings.Join(strs, " "))
	return &object.Null{}
}

// builtinTanya reads input from stdin
func builtinTanya(args ...object.Object) object.Object {
	if len(args) > 1 {
		return &object.Error{Message: "tanya() butuh maksimal 1 argumen"}
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

// builtinPanjang returns the length of string or array
func builtinPanjang(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("panjang() butuh 1 argumen, dapat %d", len(args))}
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return &object.Error{Message: fmt.Sprintf("panjang() tidak mendukung tipe %s", arg.Type())}
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
		typeName = "kosong"
	case *object.Array:
		typeName = "array"
	case *object.Map:
		typeName = "map"
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

// builtinDorong appends an element to an array
func builtinDorong(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: fmt.Sprintf("dorong() butuh 2 argumen, dapat %d", len(args))}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "argumen pertama dorong() harus array"}
	}

	newElements := make([]object.Object, len(arr.Elements)+1)
	copy(newElements, arr.Elements)
	newElements[len(arr.Elements)] = args[1]

	return &object.Array{Elements: newElements}
}

// builtinPertama returns the first element of an array
func builtinPertama(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("pertama() butuh 1 argumen, dapat %d", len(args))}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "argumen pertama() harus array"}
	}

	if len(arr.Elements) == 0 {
		return &object.Null{}
	}

	return arr.Elements[0]
}

// builtinAkhir returns the last element of an array
func builtinAkhir(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: fmt.Sprintf("akhir() butuh 1 argumen, dapat %d", len(args))}
	}

	arr, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "argumen akhir() harus array"}
	}

	if len(arr.Elements) == 0 {
		return &object.Null{}
	}

	return arr.Elements[len(arr.Elements)-1]
}
