// Package srch contains various utils to search in an array
package srch

import (
	"fmt"
)

func Partition(data []int, l int, r int) (int, error) {
	if l > r {
		return 0, fmt.Errorf("Lower bound:", l, " cannot be larger than upper bound:", r)
	}

	pivot := data[r]

	i, j := l, r - 1
	for ;i < j; {
		if data[i] > pivot && data[j] <= pivot {
			temp := data[i]
			data[i] = data[j]
			data[j] = temp
		}

		if data[i] <= pivot {
			i += 1
		}

		if data[j] > pivot {
			j -= 1
		}
	}

	data[r] = data[i]
	data[i] = pivot
	return i, nil
}
