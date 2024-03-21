// Exercise 5.19 Use panic and recover to write a function that
// contains no return statement yet returns a non-zero value.
package main

import "fmt"

func main() {
	callPaniquedAndRecovered()
	fmt.Println("We continue the execution of the program.")
}

func callPaniquedAndRecovered() {
	defer func() {
		p := recover()

		if p != nil {
			fmt.Printf("Recovered function value is %v\n", p)
		}
	}()

	panicRecover()
}

func panicRecover() {
	panic("Here is the recovered data")
}
