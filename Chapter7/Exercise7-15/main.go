// Exercise 7.15 Write a program that reads a single expression from the standart input
// prompts the user to provied values for any variables, then evaluates the expression in
// the resulting environment. Handle the errors gracefully.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Eval interface {
	Execute([]Eval) (float64, error)
}

type Expr string

func (e Expr) Execute(param []Eval) (float64, error) {
	var values []float64
	for _, par := range param {
		val, err := par.Execute(param)
		if err != nil {
			return 0, err
		}
		values = append(values, val)
	}

	switch string(e) {
	case "*":
		return values[0] * values[1], nil
	case "+":
		return values[0] + values[1], nil
	case "-":
		return values[0] - values[1], nil
	case "/":
		if values[1] == 0 {
			return 0, errors.New("unable to divide between zero")
		}
		return values[0] / values[1], nil
	default:
		return 0, errors.New("invalid operation")
	}
}

type Oper string

func (o Oper) Execute(_ []Eval) (float64, error) {
	return strconv.ParseFloat(string(o), 64)
}

func main() {
	var input Expr
	fmt.Print("Please provide a math expression: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		printError("error getting value %w", err)
		return
	}

	var par1, par2 Oper
	fmt.Print("Provide two variables, space separated to operate: ")
	_, err = fmt.Scanln(&par1, &par2)
	if err != nil {
		printError("error getting value %w", err)
		return
	}

	res, err := input.Execute([]Eval{par1, par2})
	if err != nil {
		printError("unable to compute, %w", err)
	}

	fmt.Println(res)
}

func printError(msg string, err error) {
	err = fmt.Errorf(msg, err)
	fmt.Println(err)
}
