// LinearSearch: Is a simple search algorithm that sequentially checks
// each element in a list until a match is found or the whole list
// has been searched.
// Algorithm:Call the linear search () function to search the element
// by passing the array,size and element to search in the parameter.
package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 67, 12, 10, 4, 23, 9}
	target := 4

	index := linearSearch(arr, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
