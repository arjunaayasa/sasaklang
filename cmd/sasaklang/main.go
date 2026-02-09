package main

import (
	"fmt"
	"os"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/evaluator"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/lexer"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/object"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/parser"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/repl"
)

const Version = "1.0.0"

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// Start REPL
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	switch args[0] {
	case "version", "--version", "-v":
		fmt.Printf("sasaklang versi %s\n", Version)
	case "run":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Penggunaan: sasaklang run <file>")
			os.Exit(1)
		}
		runFile(args[1])
	case "help", "--help", "-h":
		printHelp()
	default:
		// Treat as file to run (for convenience)
		if _, err := os.Stat(args[0]); err == nil {
			runFile(args[0])
		} else {
			fmt.Fprintf(os.Stderr, "Perintah tidak dikenal: %s\n", args[0])
			printHelp()
			os.Exit(1)
		}
	}
}

func runFile(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal membaca file: %s\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		fmt.Fprintln(os.Stderr, "Ada error saat parsing:")
		for _, msg := range p.Errors() {
			fmt.Fprintf(os.Stderr, "  %s\n", msg)
		}
		os.Exit(1)
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Fprintln(os.Stderr, result.Inspect())
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`SasakLang - Bahasa Pemrograman Berbasis Bahasa Sasak

Penggunaan:
  sasaklang                    Masuk ke mode REPL
  sasaklang run <file>         Jalankan file .sl
  sasaklang <file>             Jalankan file .sl (shortcut)
  sasaklang version            Tampilkan versi
  sasaklang help               Tampilkan bantuan ini

Contoh:
  sasaklang                    # Masuk REPL
  sasaklang run hello.sl       # Jalankan file
  sasaklang hello.sl           # Jalankan file (shortcut)

Dokumentasi lengkap: https://github.com/arjunaayasa/sasaklang`)
}
