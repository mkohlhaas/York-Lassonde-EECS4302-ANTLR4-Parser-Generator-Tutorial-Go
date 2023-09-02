package model

import (
	"fmt"
	"os"
)

// Variable defines a variable in our grammar.
type Variable struct {
	Name string
}

var _ Expression = &Variable{}

func (v *Variable) eval() int {
	val, exists := Program.SymbolTable[v.Name]
	if !exists {
		fmt.Printf("Variable %s is not defined.\n", v.Name)
		os.Exit(1)
	}
	return val.Value
}

func (v *Variable) String() string {
	return v.Name
}
