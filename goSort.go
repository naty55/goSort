package goSort

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
func MergeSort(matrix [][3]int32, isBigger func([3]int32, [3]int32) bool) {
	mergeSort(matrix, 0, len(matrix), isBigger)
}

func mergeSort(matrix [][3]int32, start int, end int, isBigger func([3]int32, [3]int32) bool) {
	if end <= start {
		return
	}
	p := int((end + start) / 2)
	mergeSort(matrix, start, p, isBigger)
	mergeSort(matrix, p+1, end, isBigger)
	merge(matrix, start, p, end, isBigger)

}

func merge(matrix [][3]int32, start int, middle int, end int, isBigger func([3]int32, [3]int32) bool) {
	matrix1 := append(make([][3]int32, 0), matrix[start:middle]...)
	matrix2 := append(make([][3]int32, 0), matrix[middle:end]...)
	i := 0
	j := 0
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
