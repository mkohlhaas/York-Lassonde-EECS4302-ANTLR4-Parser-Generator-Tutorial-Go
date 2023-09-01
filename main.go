package main

import (
	"fmt"
	"os"

	"expr/model"
	"expr/parser"
	"expr/visitor"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
	is, _ := os.Open(os.Args[1])
	input := antlr.NewIoStream(is)
	lexer := parser.NewExprLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewExprParser(tokens)
	prog := p.Prog()
	v := &visitor.Visitor{}
	program := v.Visit(prog).(*model.Prog)
	fmt.Println("Analyzing the input.")
	fmt.Println(program)
	fmt.Println("Running the program.\nResults:")
	for i, r := range program.Run() {
		fmt.Printf("  %d: %d\n", i+1, r)
	}
}
