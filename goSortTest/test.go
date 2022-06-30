package main

import (
	"fmt"

	"goSort.nom/goSort/goSort"
)

var matrix1 = [][3]int32{{4, 3}, {1, 2}, {1, 1}, {3, 2}, {6, 7}}

func main() {
	fmt.Println("Hello World")
	goSort.InsertionSort(matrix1, isBigger)
	fmt.Println(matrix1)
}
func isBigger(a [3]int32, b [3]int32) bool {
	return a[0] > b[0] || (a[0] == b[0] && a[1] > b[1])
}
