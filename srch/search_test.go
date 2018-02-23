package srch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	vals := []int{5, 11, 6, 99, 81, 41, 64}
	idx, err := Partition(vals, 0, len(vals) - 1)
	assert.Equal(t, err, nil)
	assert.Equal(t, idx, 4)
}
