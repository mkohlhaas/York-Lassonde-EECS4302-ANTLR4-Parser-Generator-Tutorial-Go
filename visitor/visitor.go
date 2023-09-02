// Package visitor provides all the visitor methods for the parse tree.
package visitor

import (
	"expr/model"
	"expr/parser"
	"fmt"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

// Visitor is "derived" from BaseExprVisitor.
type Visitor struct {
	parser.BaseExprVisitor
}

// Visit is a workaround to make the visitor pattern possible under Go.
func (v *Visitor) Visit(t antlr.ParseTree) any {
	return t.Accept(v)
}

var _ parser.ExprVisitor = &Visitor{}

// VisitProg loops over all variable declarations and adds them to the symbol table.
// Then it loops over all the expressions which make up the program.
// It returns the singleton `Program` for convenience.
func (v *Visitor) VisitProg(ctx *parser.ProgContext) any {
	// Declarations
	for _, decl := range ctx.AllDecl() {
		v.Visit(decl)
	}
	// Expressions
	for _, e := range ctx.AllExpr() {
		expr := v.Visit(e).(model.Expression)
		model.Program.AddExpression(expr)
	}
	return model.Program
}

// VisitDecl saves variable name with its integer value in the symbol table of the program.
func (v *Visitor) VisitDecl(ctx *parser.DeclContext) any {
	// extract new variable
	id := ctx.IDENT().GetText()
	n, _ := strconv.Atoi(ctx.NUM().GetText())

	// extract symbol
	symbol := ctx.IDENT().GetSymbol()
	line := symbol.GetLine()
	column := symbol.GetColumn() + 1

	// If var has already been declared we bail out with error message.
	variable, exists := model.Program.SymbolTable[id]
	if exists {
		fmt.Printf("Value %s at line %d, column %d has already been declared at line %d, column %d.\n", id, line, column, variable.Line, variable.Column)
		os.Exit(1)
	}
	model.Program.SymbolTable[id] = model.VarDecl{
		Value:  n,
		Line:   line,
		Column: column,
	}
	return nil
}

// VisitAddition visits an AdditionContext and creates an Addition AST.
func (v *Visitor) VisitAddition(ctx *parser.AdditionContext) any {
	left := v.Visit(ctx.Expr(0)).(model.Expression)
	right := v.Visit(ctx.Expr(1)).(model.Expression)
	return &model.Addition{Left: left, Right: right}
}

// VisitMultiplication visits a MultiplicationContext and creates a Multiplication AST.
func (v *Visitor) VisitMultiplication(ctx *parser.MultiplicationContext) any {
	left := v.Visit(ctx.Expr(0)).(model.Expression)
	right := v.Visit(ctx.Expr(1)).(model.Expression)
	return &model.Multiplication{Left: left, Right: right}
}

// VisitVariable visits a VariableContext and creates a Variable AST.
func (v *Visitor) VisitVariable(ctx *parser.VariableContext) any {
	return &model.Variable{Name: ctx.IDENT().GetText()}
}

// VisitNumber visits a NumberContext and creates a Number AST.
func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	n, _ := strconv.Atoi(ctx.NUM().GetText())
	return &model.Number{Num: n}
}
