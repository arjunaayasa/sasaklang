package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/evaluator"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/lexer"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/object"
	"github.com/arjunaayasa/sasaklang/pkg/sasaklang/parser"
)

const PROMPT = "sasak>> "

const LOGO = `
   _____                 __   __                   
  / ___/____ _________ _/ /__/ /   ____ _____  ____ _
  \__ \/ __ '/ ___/ __ '/ //_/ /   / __ '/ __ \/ __ '/
 ___/ / /_/ (__  ) /_/ / ,< / /___/ /_/ / / / / /_/ / 
/____/\__,_/____/\__,_/_/|_/_____/\__,_/_/ /_/\__, /  
                                             /____/   
`

// Start starts the REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Fprint(out, LOGO)
	fmt.Fprintln(out, "Selamat datang di SasakLang REPL!")
	fmt.Fprintln(out, "Ketik 'exit' atau tekan Ctrl+D untuk keluar.")
	fmt.Fprintln(out)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Fprintln(out, "\nSampai jumpa!")
			return
		}

		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == "exit" || line == "keluar" {
			fmt.Fprintln(out, "Sampai jumpa!")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			// Don't print null for expression statements
			if evaluated.Type() != object.NULL_OBJ {
				fmt.Fprintln(out, evaluated.Inspect())
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	fmt.Fprintln(out, "Ada error saat parsing:")
	for _, msg := range errors {
		fmt.Fprintf(out, "  %s\n", msg)
	}
}
