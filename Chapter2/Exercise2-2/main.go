// Write a general-purpose unit-conversion program, analogous to cf that reads numbers from its command-line
// arguments or from standart input if there are not arguments, and converts
// each number into units like temp in Celcius and Fahrenheit,
// Lenght in feet and meters, weigth in pounds and kilograms, and the like.
package main

import (
	"Exercise2-2/generalconverter"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]

	if len(input) == 0 {
		inputData := bufio.NewScanner(os.Stdin)
		for inputData.Scan() {
			convertDataToDifferentUnits(inputData.Text())
		}
	} else {
		for _, data := range input {
			convertDataToDifferentUnits(data)
		}
	}
}

func convertDataToDifferentUnits(input string) {
	value, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Unable to convert input: %s - %v\n", input, err)
		return
	}

	c := generalconverter.Celsius(value)
	f := generalconverter.Fahrenheit(value)
	m := generalconverter.Meters(value)
	ft := generalconverter.Feets(value)
	kg := generalconverter.Kilogams(value)
	pd := generalconverter.Pounds(value)

	fmt.Printf("%s\t=\t%5s\n", c, c.ToF())
	fmt.Printf("%s\t=\t%s\n", f, f.ToC())
	fmt.Printf("%s\t=\t%s\n", m, m.ToFeet())
	fmt.Printf("%s\t=\t%s\n", ft, ft.ToM())
	fmt.Printf("%s\t=\t%s\n", kg, kg.ToPounds())
	fmt.Printf("%s\t=\t%s\n", pd, pd.ToKg())
}
