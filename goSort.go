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
// Number of Algorithms so far :  8
//
// NOTE:  This package is an eductional project for learning and practicing algorithmic thinking and NOT intended to be used outside the classroom
package goSort

func BubbleSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	for j := len(matrix); j > 1; j-- {
		for i := 1; i < j; i++ {
			if isBigger(matrix[i-1], matrix[i]) {
				swap(matrix, i-1, i)
			}
		}
	}

}

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
		// Insert Element in idx'th place
		for i := key - 1; i >= idx; i-- {
			matrix[i+1] = matrix[i]
		}
		matrix[idx] = current

	}
}

func HeapSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	heapifyMatrix(matrix, isBigger)
	for end := len(matrix) - 1; end >= 0; end-- {
		swap(matrix, end, 0)
		maxHeapify(matrix, 0, end, isBigger)
	}
}

func heapifyMatrix(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	for i := int(len(matrix) / 2); i >= 0; i-- {
		maxHeapify(matrix, i, len(matrix), isBigger)
	}
}

func maxHeapify(heap [][3]int32, i int, heapEnd int, isBigger func([3]int32, [3]int32) bool) {
	left := 2 * i
	right := 2*i + 1
	largest := i
	if left < heapEnd && isBigger(heap[left], heap[largest]) {
		largest = left
	}
	if right < heapEnd && isBigger(heap[right], heap[largest]) {
		largest = right
	}
	if largest != i {
		swap(heap, largest, i)
		maxHeapify(heap, largest, heapEnd, isBigger)
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

func QuickSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	quickSort(matrix, 0, len(matrix)-1, isBigger)
}

func quickSort(matrix [][3]int32, start int, end int, isBigger func([3]int32, [3]int32) bool) {
	if end <= start || start < 0 {
		return
	}
	p := partitionLast(matrix, start, end, isBigger)
	quickSort(matrix, p+1, end, isBigger)
	quickSort(matrix, start, p-1, isBigger)
}

func partitionLast(matrix [][3]int32, start int, end int, isBigger func([3]int32, [3]int32) bool) int {
	pivot := matrix[end]
	i := start
	for j := start; j < end; j++ {
		if isBigger(pivot, matrix[j]) {
			swap(matrix, i, j)
			i++
		}
	}
	swap(matrix, i, end)
	return i
}

func swap(matrix [][3]int32, i int, j int) {
	temp := matrix[i]
	matrix[i] = matrix[j]
	matrix[j] = temp
}

func RadixSort(matrix [][3]int32) {
	for i := 2; i >= 0; i-- {
		countSort(matrix, i)
	}
}

func countSort(matrix [][3]int32, k int) {

}

func SelctionSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	for i := 0; i < len(matrix); i++ {
		minIdx := i
		for j := i + 1; j < len(matrix); j++ {
			if isBigger(matrix[minIdx], matrix[j]) {
				minIdx = j
			}
		}
		swap(matrix, i, minIdx)
	}
}

func StoogeSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	stoogeSort(matrix, 0, len(matrix)-1, isBigger)
}

func stoogeSort(matrix [][3]int32, i int, j int, isBigger func([3]int32, [3]int32) bool) {
	if len(matrix) <= 1 {
		return
	}
	if isBigger(matrix[i], matrix[j]) {
		swap(matrix, i, j)
	}
	if i+1 >= j {
		return
	}
	k := int((j - i + 1) / 3)
	stoogeSort(matrix, i, j-k, isBigger)
	stoogeSort(matrix, i+k, j, isBigger)
	stoogeSort(matrix, i, j-k, isBigger)
}
