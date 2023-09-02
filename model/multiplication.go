package model

import "fmt"

// Multiplication defines a multiplication in our grammar.
type Multiplication struct {
	Left, Right Expression
}

var _ Expression = &Multiplication{}

func (m *Multiplication) eval() int {
	return m.Left.eval() * m.Right.eval()
}

func (m *Multiplication) String() string {
	return fmt.Sprintf("%s * %s", m.Left, m.Right)
}
