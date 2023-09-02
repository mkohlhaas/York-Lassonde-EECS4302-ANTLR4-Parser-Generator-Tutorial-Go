package model

import (
	"fmt"
	"strings"
)

var Program = &Prog{SymbolTable: make(SymbolTable)} // Program is a singleton.

// VarDecl defines a variable declaration in our grammar.
type VarDecl struct {
	Value  int
	Line   int
	Column int
}

func (v VarDecl) String() string {
	return fmt.Sprintf("[value: %d, line: %d, column: %d]", v.Value, v.Line, v.Column)
}

// SymbolTable reflectsa a symbol table. Holds name of the variable and its declaration details.
type SymbolTable map[string]VarDecl

// Prog in our grammar is a list of expressions and a symbol table.
type Prog struct {
	Expressions []Expression // List of all the expressions our program is made of.
	SymbolTable SymbolTable  // Stores identifiers/variables and their values.
}

// AddExpression adds an expression to our program.
func (p *Prog) AddExpression(e Expression) {
	p.Expressions = append(p.Expressions, e)
}

// Run runs our program, i.e. all the expressions a program consists of.
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
		fmt.Fprintf(b, "  %s: %s\n", k, v)
	}
	fmt.Fprintln(b, "\nExpressions:")
	for i, e := range p.Expressions {
		fmt.Fprintf(b, "  %d: %s\n", i+1, e)
	}
	return b.String()
}
