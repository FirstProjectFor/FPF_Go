package algorithm

func duplicateZeros(arr []int) {
	arrLength := len(arr)
	if arrLength == 0 {
		return
	}
	arrCopy := copyArr(arr)

	index := 0
	for _, v := range arrCopy {
		arr[index] = v
		index++
		if v == 0 && index < arrLength {
			arr[index] = 0
			index++
		}
		if index >= arrLength {
			break
		}
	}
}

func copyArr(arr []int) []int {
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	return arrCopy
}
