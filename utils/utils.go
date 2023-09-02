// Package utils provides one public function - RunProgram() - to be used by main().
package utils

import (
	"errors"
	"fmt"
	"io"
	"os"

	el "expr/error_listener"
	m "expr/model"
	ep "expr/parser"
	v "expr/visitor"

	"github.com/antlr4-go/antlr/v4"
	"github.com/duke-git/lancet/v2/tuple"
)

// RunProgram generates a program with and Antlr4 parser and then runs it.
func RunProgram() {
	parser := getParser()
	program := parseTree2Model(parser)
	fmt.Println(program)
	printResults(program)
}

// Typical procedure to get the Antlr4 system going.
func getParser() *ep.ExprParser {
	input := antlr.NewIoStream(getInputStream())
	lexer := ep.NewExprLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := ep.NewExprParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(new(el.ExprErrorListener))
	return parser
}

// Parses command line arguements to decide which input stream to take - standard in or file.
func getInputStream() io.Reader {
	is := os.Stdin
	var err error
	switch len(os.Args) {
	case 1:
		fmt.Println("Reading from standard input.")
	case 2:
		filename := os.Args[1]
		is, err = os.Open(filename)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("File %s does not exist.\n", filename)
			os.Exit(1)
		}
		fmt.Printf("Reading from file: %s\n", filename)
	default:
		fmt.Printf("Usage: %s [filename]\n", os.Args[0])
		fmt.Printf("Without a [filename] it will read from standard input.\n")
		os.Exit(1)
	}
	return is
}

// Generates model by using visitor and parser.
func parseTree2Model(p *ep.ExprParser) *m.Prog {
	v := new(v.Visitor)
	return v.Visit(p.Prog()).(*m.Prog)
}

func printResults(program *m.Prog) {
	for i, t := range tuple.Zip2(program.Expressions, program.Run()) {
		fmt.Printf("%d: %v => %v\n", i+1, t.FieldA, t.FieldB)
	}
}
