package model

import (
	"strconv"
)

// Number defines a number in our grammar.
type Number struct {
	Num int
}

var _ Expression = &Number{}

func (n *Number) eval() int {
	return n.Num
}

func (n *Number) String() string {
	return strconv.Itoa(n.Num)
}
