// Exercise 7.7 Explain why the help message contains °C when the default value 20.0 does not.
package main

import (
	"flag"
	"fmt"
)

func CelsiusFlag(name string, value Celsius, usage string) *celsiusFlag {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f
}

// The helper uses the fmt.stringer helper to parse the result, check the return of the CelsiusFlag func
func (f celsiusFlag) String() string {
	return fmt.Sprintf("%g° Celsius", f.Celsius)
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
