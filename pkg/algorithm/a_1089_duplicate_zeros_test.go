package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	var arr, arrDuplicateZero []int

	arr = []int{1, 2}
	arrDuplicateZero = []int{1, 2}
	duplicateZeros(arr)
	assert.Equal(t, arr, arrDuplicateZero)

	arr = []int{1, 0, 1}
	arrDuplicateZero = []int{1, 0, 0}
	duplicateZeros(arr)
	assert.Equal(t, arr, arrDuplicateZero)

	arr = []int{1, 0, 0, 1, 1, 1, 1, 1}
	arrDuplicateZero = []int{1, 0, 0, 0, 0, 1, 1, 1}
	duplicateZeros(arr)
	assert.Equal(t, arr, arrDuplicateZero)

	arr = []int{1, 0, 0}
	arrDuplicateZero = []int{1, 0, 0}
	duplicateZeros(arr)
	assert.Equal(t, arr, arrDuplicateZero)
}
