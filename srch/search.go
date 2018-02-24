// Package srch contains various utils to search in an array
package srch

import (
	"fmt"
)

func Partition(data []int, l int, r int) (int, error) {
	if l > r {
		return 0, fmt.Errorf("Lower bound: %d cannot be larger than upper bound: %d", l, r)
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


func FindSmallest(data []int, l int, r int, k int) error {
	fmt.Println(l, r, k)
	if k > len(data) {
		return fmt.Errorf("k cannot be larger than length of data")
	}

	pivot, err := Partition(data, l, r)
	if err != nil {
		return err
	}

	if pivot > l + k - 1 {
		return FindSmallest(data, l, pivot - 1, k)
	}

	if pivot == l + k - 1 {
		return nil
	}

	return FindSmallest(data, pivot + 1, r, k - pivot - 1)
}
