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
