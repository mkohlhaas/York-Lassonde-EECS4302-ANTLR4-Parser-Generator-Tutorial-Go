package model

import (
	"fmt"
	"strings"
)

var Program = &Prog{SymbolTable: make(SymbolTable)} // Program is a singleton.

type SymbolTable map[string]int

type Prog struct {
	Expressions []Expression // List of all the expressions our program is made of.
	SymbolTable SymbolTable  // Stores identifiers/variables and their values.
}

func (p *Prog) AddExpression(e Expression) {
	p.Expressions = append(p.Expressions, e)
}

func (p *Prog) Run() (result []int) {
	for _, e := range p.Expressions {
		result = append(result, e.eval())
	}
	return
}

func (p *Prog) String() string {
	b := &strings.Builder{}
	fmt.Println("Variables:")
	for k, v := range p.SymbolTable {
		fmt.Printf("  %s: %d\n", k, v)
	}
	fmt.Fprintln(b, "\nExpressions:")
	for i, e := range p.Expressions {
		fmt.Fprintf(b, "  %d: %s\n", i+1, e)
	}
	return b.String()
}
