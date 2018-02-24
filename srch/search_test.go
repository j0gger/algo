package srch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	vals := []int{5, 11, 6, 99, 81, 41, 64}
	idx, err := Partition(vals, 0, len(vals) - 1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 4, idx)
}


func TestFindSmallest(t *testing.T) {
	vals := []int{55, 11, 10, 34, 14, 12, 26}
	err := FindSmallest(vals, 0, len(vals) - 1, 6)

	assert.Equal(t, nil, err)
	assert.Equal(t, []int{12, 11, 10, 14, 26, 34}, vals[:6])
}



func TestFindSmallest2(t *testing.T) {
	vals := []int{3, 2, 1}
	err := FindSmallest(vals, 0, len(vals) - 1, 2)

	assert.Equal(t, nil, err)
	assert.Equal(t, []int{1, 2}, vals[:2])
}

