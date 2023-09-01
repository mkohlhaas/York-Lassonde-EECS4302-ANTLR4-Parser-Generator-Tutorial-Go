package model

import (
	"strconv"
)

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
