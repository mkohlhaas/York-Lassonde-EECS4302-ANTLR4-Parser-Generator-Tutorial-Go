package model

import "fmt"

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
