package model

// Expression can be evaluated.
type Expression interface {
	eval() int
}
