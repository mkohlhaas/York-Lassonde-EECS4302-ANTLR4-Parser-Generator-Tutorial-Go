package visitor

import (
	"expr/model"
	"expr/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.BaseExprVisitor
}

func (v *Visitor) Visit(t antlr.ParseTree) any {
	return t.Accept(v)
}

var _ parser.ExprVisitor = &Visitor{}

// VisitProg loops over all variable declarations and adds them to the symbol table.
// Then loops over all the expressions which make up the program.
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
	n, _ := strconv.Atoi(ctx.NUM().GetText())
	model.Program.SymbolTable[ctx.IDENT().GetText()] = n
	return nil
}

func (v *Visitor) VisitAddition(ctx *parser.AdditionContext) any {
	left := v.Visit(ctx.Expr(0)).(model.Expression)
	right := v.Visit(ctx.Expr(1)).(model.Expression)
	return &model.Addition{Left: left, Right: right}
}

func (v *Visitor) VisitMultiplication(ctx *parser.MultiplicationContext) any {
	left := v.Visit(ctx.Expr(0)).(model.Expression)
	right := v.Visit(ctx.Expr(1)).(model.Expression)
	return &model.Multiplication{Left: left, Right: right}
}

func (v *Visitor) VisitVariable(ctx *parser.VariableContext) any {
	return &model.Variable{Name: ctx.IDENT().GetText()}
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) any {
	n, _ := strconv.Atoi(ctx.NUM().GetText())
	return &model.Number{Num: n}
}
