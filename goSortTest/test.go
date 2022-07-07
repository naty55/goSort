package main

import (
	"fmt"

	"goSort.nom/goSort"
)

var matrix1 = [][3]int32{{4, 3}, {1, 2}, {1, 1}, {3, 2}, {6, 7}, {4, 4}}  // Random Matrix
var matrix2 = [][3]int32{{6}, {5}, {5, 1}, {4}, {3}, {2}, {1}, {0}, {-2}} // Complete reversed Matrix
var matrix3 = [][3]int32{{6}, {6}, {4}, {5}, {3}, {3}, {3}, {2}, {4}}     // Matrix with duplicates
var matrix4 = [][3]int32{{1}, {2}, {3}, {3}, {4}, {5}, {6}, {7}}          // Sorted Matrix

var matrices = [][][3]int32{matrix1, matrix2, matrix3, matrix4}

func main() {
	testSort("BubbleSort", goSort.BubbleSort, isBigger)
	testSort("HeapSort", goSort.HeapSort, isBigger)
	testSort("InsertionSort", goSort.InsertionSort, isBigger)
	testSort("InsertionBinarySort", goSort.BinaryInsertionSort, isBigger)
	testSort("MergeSort", goSort.MergeSort, isBigger)
	testSort("QuickSort", goSort.QuickSort, isBigger)
	testSort("SelectionSort", goSort.SelctionSort, isBigger)
}
func isBigger(a [3]int32, b [3]int32) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}

func testSort(name string, sortAlg func([][3]int32, func([3]int32, [3]int32) bool), isBigger func([3]int32, [3]int32) bool) {
	fmt.Printf("%s Algorithm : ", name)
	failed := false
	var failed_test [][3]int32
	for i := 0; i < len(matrices) && !failed; i++ {
		test_matrix := append(make([][3]int32, 0), matrices[i]...)
		sortAlg(test_matrix, isBigger)
		if !isSorted(test_matrix, isBigger) {
			failed = true
			failed_test = test_matrix
			break
		}
	}

	if failed {
		fmt.Printf("Didn't Pass the test result: ")
		fmt.Println(failed_test)
		panic("Error. didn't pass the test, result ")
	} else {
		fmt.Println("Passed the test")
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
