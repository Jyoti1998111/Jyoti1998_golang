// Golang program to print Fibonacci series
// using "for" loop.
//Fibonacci : fibonacci is the sequence of numbers in which each
//number in the sequence is equal to the sum of two numbers before it.

//Algorithm : we will read an integer number and print the fibonacci
//series. the source code to print the fibonacci series using the for loop

package main

import "fmt"

func main() {
	var num1 int = 0
	var num2 int = 1
	var num3 int = 2
	var num4 int = 3

	fmt.Printf("%d", num1)
	fmt.Printf(" %d", num2)

	for count := 1; count <= 8; count++ {
		num3 = num1 + num2
		fmt.Printf(" %d", num3)

		num1 = num2
		num2 = num3

		num4 = num4 + 1
	}
}
