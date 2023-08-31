// Main entry point for tutorial
package main

import (
	"fmt"
	"os"

	"expr/parser"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	is, _ := os.Open(os.Args[1])
	input := antlr.NewIoStream(is)
	lexer := parser.NewExprLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(tokens)
	prog := p.Prog()
	fmt.Printf("%v\n", prog)
}
