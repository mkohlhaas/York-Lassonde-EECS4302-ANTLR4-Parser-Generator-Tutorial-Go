// Package parser consists of the Antlr4 generated functions and some helper methods.
package parser

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

// ExprErrorListener will be called on syntax errors, print an error message and bail out.
type ExprErrorListener struct {
	*antlr.DefaultErrorListener
}

// SyntaxError will be called by the parser in case of syntax errors. Prints error message and exits.
func (*ExprErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	fmt.Fprintf(os.Stderr, "Syntax error at line %d, column %d: %s\n", line, column+1, msg)
	os.Exit(1)
}
