package algorithm

func decrypt(code []int, k int) []int {

	if code == nil || len(code) == 1 {
		return code
	}

	if k == 0 {
		return make([]int, len(code))
	}

	result := make([]int, len(code))

	return result
}
