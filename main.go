package main

import (
	"bubbleSort/bubbleSort"
	"fmt"
)

func main() {
	bubbleSortResponse := bubbleSort.BubbleSort([]int{3, 2, 6, 2, 9, 23, 67, 4, 1, -5, 0})
	fmt.Printf("Response after bubble Sort : ")
	fmt.Println(bubbleSortResponse)
}
