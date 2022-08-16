package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Interpret(context string) bool
}
type TerminalExpression struct {
	Data string
}

func NewTerminalExpression(data string) *TerminalExpression {
	return &TerminalExpression{Data: data}
}

func (t TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.Data) {
		return true
	}

	return false
}

type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewOrExpression(expr1 Expression, expr2 Expression) *OrExpression {
	return &OrExpression{Expr1: expr1, Expr2: expr2}
}

func (o OrExpression) Interpret(context string) bool {
	return o.Expr1.Interpret(context) || o.Expr2.Interpret(context)
}

type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

func NewAndExpression(expr1 Expression, expr2 Expression) *AndExpression {
	return &AndExpression{Expr1: expr1, Expr2: expr2}
}

func (a AndExpression) Interpret(context string) bool {
	return a.Expr1.Interpret(context) && a.Expr2.Interpret(context)
}

func main() {
	lee := NewTerminalExpression("Lee")
	wang := NewTerminalExpression("Wang")
	oe := NewOrExpression(lee, wang)
	b := oe.Interpret("Lee")
	fmt.Printf("b: %v\n", b)

	yang := NewTerminalExpression("yang")
	married := NewTerminalExpression("Married")
	ae := NewAndExpression(yang, married)
	b2 := ae.Interpret("Married yang")
	fmt.Printf("b2: %v\n", b2)
}
