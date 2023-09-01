package model

type Variable struct {
	Name string
}

var _ Expression = &Variable{}

func (v *Variable) eval() int {
	val, ok := Program.SymbolTable[v.Name]
	if ok {
		return val
	} else {
		panic(ok)
	}
}

func (v *Variable) String() string {
	return v.Name
}
