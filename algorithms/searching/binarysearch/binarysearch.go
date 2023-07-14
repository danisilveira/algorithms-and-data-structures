package binarysearch

import (
	"golang.org/x/exp/constraints"
)

func Search[T constraints.Ordered](items []T, search T) (int, int) {
	index := 0
	steps := 0

	if len(items) == 0 {
		steps++
		return -1, steps
	}

	middle := len(items) / 2

	if search > items[middle] {
		index, steps = Search(items[middle+1:], search)
		if index >= 0 {
			index += middle + 1
		}
	}

	if search < items[middle] {
		index, steps = Search(items[:middle], search)
	}

	if search == items[middle] {
		index = middle
	}

	steps++
	return index, steps
}
