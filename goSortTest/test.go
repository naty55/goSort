package main

import (
	"fmt"

	"goSort.nom/goSort"
)

var matrix = [][3]int32{{4, 3}, {1, 2}, {1, 1}, {3, 2}, {6, 7}, {4, 4}}

func main() {
	testSort("InsertionSort", goSort.InsertionSort, isBigger)
	testSort("MergeSort", goSort.MergeSort, isBigger)
}
func isBigger(a [3]int32, b [3]int32) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}

func testSort(name string, sortAlg func([][3]int32, func([3]int32, [3]int32) bool), isBigger func([3]int32, [3]int32) bool) {
	fmt.Printf("%s Algorithm : ", name)
	test_matrix := append(make([][3]int32, 0), matrix...)
	sortAlg(test_matrix, isBigger)
	if isSorted(test_matrix, isBigger) {
		fmt.Println("Passed the test")
	} else {
		fmt.Printf("Didn't Pass the test result: ")
		fmt.Println(test_matrix)
		panic("Error. didn't pass the test, result ")
	}
}

func isSorted(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) bool {
	for i := 1; i < len(matrix); i++ {
		if isBigger(matrix[i-1], matrix[i]) {
			return false
		}
	}
	return true
}
