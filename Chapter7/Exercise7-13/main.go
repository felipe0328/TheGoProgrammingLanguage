// Exercise 7.13 Add a String method to Expr to pretty-print the syntax tree. Check that the results,
// when parsed again, yield and equivalent tree.
package main

import (
	"fmt"
	"math"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	var arguments []string

	for _, item := range c.args {
		arguments = append(arguments, item.String())
	}

	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(arguments, ","))
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
