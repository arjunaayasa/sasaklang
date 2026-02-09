package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/ast"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	INTEGER_OBJ      ObjectType = "INTEGER"
	STRING_OBJ       ObjectType = "STRING"
	BOOLEAN_OBJ      ObjectType = "BOOLEAN"
	NULL_OBJ         ObjectType = "NULL"
	RETURN_VALUE_OBJ ObjectType = "RETURN_VALUE"
	ERROR_OBJ        ObjectType = "ERROR"
	FUNCTION_OBJ     ObjectType = "FUNCTION"
	BUILTIN_OBJ      ObjectType = "BUILTIN"
	ARRAY_OBJ        ObjectType = "ARRAY"
	MAP_OBJ          ObjectType = "MAP"
)

// Object is the interface all objects implement
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Hashable is implemented by objects that can be map keys
type Hashable interface {
	HashKey() HashKey
}

// HashKey represents a hash key for maps
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// Integer represents an integer value
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	if b.Value {
		return "bener"
	}
	return "salah"
}
func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

// Null represents null (kosong)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "kosong" }

// ReturnValue wraps a return value
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// Error represents a runtime error
type Error struct {
	Message string
	Line    int
	Column  int
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	if e.Line > 0 {
		return fmt.Sprintf("Error [baris %d, kolom %d]: %s", e.Line, e.Column, e.Message)
	}
	return fmt.Sprintf("Error: %s", e.Message)
}

// Function represents a function object
type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("pungsi")
	if f.Name != "" {
		out.WriteString(" " + f.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") { ... }")
	return out.String()
}

// BuiltinFunction is the type for builtin functions
type BuiltinFunction func(args ...Object) Object

// Builtin represents a builtin function
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "fungsi bawaan" }

// Array represents an array
type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// MapPair represents a key-value pair in a map
type MapPair struct {
	Key   Object
	Value Object
}

// Map represents a map/dictionary
type Map struct {
	Pairs map[HashKey]MapPair
}

func (m *Map) Type() ObjectType { return MAP_OBJ }
func (m *Map) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// Environment stores variable bindings
type Environment struct {
	store  map[string]Object
	consts map[string]bool // tracks which variables are constants
	outer  *Environment
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	return &Environment{
		store:  make(map[string]Object),
		consts: make(map[string]bool),
		outer:  nil,
	}
}

// NewEnclosedEnvironment creates a new environment with an outer scope
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a variable from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return obj, ok
}

// Set sets a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// SetConst sets a constant in the environment
func (e *Environment) SetConst(name string, val Object) Object {
	e.store[name] = val
	e.consts[name] = true
	return val
}

// IsConst checks if a variable is a constant
func (e *Environment) IsConst(name string) bool {
	if isConst, ok := e.consts[name]; ok && isConst {
		return true
	}
	if e.outer != nil {
		return e.outer.IsConst(name)
	}
	return false
}

// Update updates an existing variable (for assignment)
func (e *Environment) Update(name string, val Object) (Object, bool) {
	if _, ok := e.store[name]; ok {
		e.store[name] = val
		return val, true
	}
	if e.outer != nil {
		return e.outer.Update(name, val)
	}
	return nil, false
}
