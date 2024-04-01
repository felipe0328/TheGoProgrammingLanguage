// Exercise 7.14 Define a new concrete type that satisfies Expr interface and provides a new operation
// such as computing the minimum value of its operands. Since the Parse function does not create
// instances of this new type, to use it you will need to construct a syntax tree directly
// (or extend the parser).
package main

import (
	"fmt"
	"math"
)

type minimum struct {
	x float64
	b binary
}

func (m minimum) String() string {
	return m.b.String()
}
func (m minimum) Check(vars map[Var]bool) error { return nil }
func (m minimum) Eval(e Env) float64 {
	x := m.b.x.Eval(e)
	y := m.b.y.Eval(e)

	if x < y {
		m.x = x
	} else {
		m.x = y
	}

	fmt.Println("The minimum operand is: ", m.x)

	return m.b.Eval(e)
}

func main() {
	mathFunc := "sqrt(A/pi)"
	expr, err := Parse(mathFunc)
	if err != nil {
		fmt.Println("Error parsing exp: ", err)
	}
	result := expr.Eval(Env{"A": 87616, "pi": math.Pi})
	fmt.Printf("Exp: %s\n", expr)
	fmt.Printf("%s => %f\n", mathFunc, result)
}
