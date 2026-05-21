package main

import (
	"fmt"
	"sortAlgos/bubbleSort"
	"sortAlgos/insertionSort"
	"sortAlgos/mergeSort"
	"sortAlgos/selectionSort"
)

func main() {
	arr := []int{3, 2, 6, 2, 9, 23, 67, 4, 1, -5, 0}
	bubbleSortResponse := bubbleSort.BubbleSort(append([]int(nil), arr...))
	fmt.Printf("Response after bubble Sort    : ")
	fmt.Println(bubbleSortResponse)

	selectionSortResponse := selectionSort.SelectionSort(append([]int(nil), arr...))
	fmt.Printf("Response after Selection Sort : ")
	fmt.Println(selectionSortResponse)

	insertionSortResponse := insertionSort.InsertionSort(append([]int(nil), arr...))
	fmt.Printf("Response for Insertion Sort   : ")
	fmt.Println(insertionSortResponse)

	mergeSortResponse := mergeSort.MergeSort(append([]int(nil), arr...))
	fmt.Printf("Response for Merge Sort       : ")
	fmt.Println(mergeSortResponse)
}
