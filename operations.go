// Impliment basic math operations on two numbers.
// operator : An operator is a symbol that performs operations on a
// value.
package main

import "fmt"

func main() {
	num1 := 6
	num2 := 2
	sum := num1 + num2
	fmt.Printf("%d+%d=%d\n", num1, num2, sum)

	difference := num1 - num2
	fmt.Printf("%d-%d=%d\n", num1, num2, difference)

	multiply := num1 * num2
	fmt.Printf("%d*%d=%d\n", num1, num2, multiply)

	divide := num1 / num2
	fmt.Printf("%d/%d=%d\n", num1, num2, divide)

}
