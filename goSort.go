// Package goSort provides an implementation for many types of SORT ALGORITHMS.
//
// To Make things more fun, the implementation will be for [][3]int32 slices.
//
// ===============
//
// So who says which slice is "bigger"?
//
// The defnition for which slice is bigger than others has been left for the user, the user will have to pass a comparsion function taking 2 [3]int32 arrays and return @true if and only if the first is the bigger (according to his/her definition)
//
// Author: Naty Mina
//
// NOTE:  This package is an eductional project for learning and practicing algorithmic thinking and NOT intended to be used outside the classroom
package goSort

// === Clasic InsertionSort ===
//
func InsertionSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	for key := 1; key < len(matrix); key++ {
		current := matrix[key]
		var idx int
		for idx = key - 1; idx >= 0; idx-- {
			if isBigger(matrix[idx], current) {
				matrix[idx+1] = matrix[idx]
			} else {
				break
			}
		}
		matrix[idx+1] = current

	}
}
func BinaryInsertionSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	for key := 1; key < len(matrix); key++ {
		current := matrix[key]
		var idx int
		// Binary Search to insert current element
		for start, end := 0, key; start < end; {
			middle := int((end + start) / 2)
			idx = middle
			if isBigger(matrix[middle], current) {
				end = middle
			} else if isBigger(current, matrix[middle]) {
				idx++
				start = middle + 1
			} else {
				break
			}
		}
		// Insert Element in place
		for i := key - 1; i >= idx; i-- {
			matrix[i+1] = matrix[i]
		}
		matrix[idx] = current

	}
}

func MergeSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	mergeSort(matrix, 0, len(matrix), isBigger)
}

func mergeSort(matrix [][3]int32, start int, end int, isBigger func([3]int32, [3]int32) bool) {
	if end <= start+1 {
		return
	}
	middle := int((end + start) / 2)
	mergeSort(matrix, start, middle, isBigger)
	mergeSort(matrix, middle, end, isBigger)
	merge(matrix, start, middle, end, isBigger)

}

func merge(matrix [][3]int32, start int, middle int, end int, isBigger func([3]int32, [3]int32) bool) {
	matrix1 := append(make([][3]int32, 0), matrix[start:middle]...)
	matrix2 := append(make([][3]int32, 0), matrix[middle:end]...)
	i, j := 0, 0
	var key int
	for key = start; i < len(matrix1) && j < len(matrix2); key++ {
		if isBigger(matrix1[i], matrix2[j]) {
			matrix[key] = matrix2[j]
			j++
		} else {
			matrix[key] = matrix1[i]
			i++
		}
	}

	for ; i < len(matrix1); i++ {
		matrix[key] = matrix1[i]
		key++
	}

	for ; j < len(matrix2); j++ {
		matrix[key] = matrix2[j]
		key++
	}
}
