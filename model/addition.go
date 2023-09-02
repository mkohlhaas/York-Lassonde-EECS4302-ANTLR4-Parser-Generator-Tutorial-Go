// Package model is the AST for a program as defined in the grammar Expr.g4.
package model

import "fmt"

// Addition defines addition in our grammar.
type Addition struct {
	Left, Right Expression
}

var _ Expression = &Addition{}

func (a *Addition) eval() int {
	return a.Left.eval() + a.Right.eval()
}

func (a *Addition) String() string {
	return fmt.Sprintf("%s + %s", a.Left, a.Right)
}
