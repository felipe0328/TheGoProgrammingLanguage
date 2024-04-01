package main

import (
	"fmt"
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
