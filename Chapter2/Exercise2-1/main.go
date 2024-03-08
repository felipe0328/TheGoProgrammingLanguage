// Add types, constant and funtions to tempconv for procesing temperatures in Kelvin scale
// Where Kelving is -273.15°C and a difference of 1K has the same magnitude as 1°C
package main

import (
	"Exercise2-1/tempconv"
	"fmt"
)

func main() {
	celsius := tempconv.Celsius(50)
	fahrenheit := tempconv.Fahrenheit(50)
	kelvin := tempconv.Kelvin(50)

	fmt.Printf("%s are %s\n", celsius, tempconv.CToF(celsius))
	fmt.Printf("%s are %s\n", fahrenheit, tempconv.FToC(fahrenheit))
	fmt.Printf("%s are %s\n", kelvin, tempconv.KToC(kelvin))
}
