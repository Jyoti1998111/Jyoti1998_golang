// Golang program to Bubble sort
//Bubble sort: Bubble sort is a basic algorithm for arranging a string
//of numbers or other elements in the correct order.
//Algorithm: traverse from  left and compare adjacent elements and
//the higher one is placed at right side.
//In this way, the largest element is moved to the rightmost end at
//first.

package main

import "fmt"

func BubbleSort(items []int) {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				temp := items[j]
				items[j] = items[j+1]
				items[j+1] = temp
			}
		}
	}
}
func main() {
	items := []int{12, 3, 7, 34, 57, 67, 89, 14, 21}
	fmt.Println(items)
	BubbleSort(items)
	fmt.Println(items)

}
