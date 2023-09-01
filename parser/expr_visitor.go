// Code generated from Expr.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Expr

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ExprParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by ExprParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by ExprParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by ExprParser#multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by ExprParser#addition.
	VisitAddition(ctx *AdditionContext) interface{}
}
