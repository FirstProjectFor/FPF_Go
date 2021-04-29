package algorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecrypt(t *testing.T) {
	result := decrypt(nil, 0)
	assert.Nil(t, result)

	result = decrypt([]int{}, 0)
	assert.Equal(t, 0, len(result))

	result = decrypt([]int{1, 2}, 0)
	assert.Equal(t, 0, result[0])
	assert.Equal(t, 0, result[1])

	result = decrypt([]int{5, 7, 1, 4}, 3)
	assert.Equal(t, 12, result[0])
	assert.Equal(t, 10, result[1])
	assert.Equal(t, 16, result[2])
	assert.Equal(t, 13, result[3])

	result = decrypt([]int{2, 4, 9, 3}, -2)
	assert.Equal(t, 12, result[0])
	assert.Equal(t, 5, result[1])
	assert.Equal(t, 6, result[2])
	assert.Equal(t, 13, result[3])
}
