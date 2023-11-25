package linearsearch

import "cmp"

func Search[S ~[]T, T cmp.Ordered](items S, search T) int {
	for index, item := range items {
		if item == search {
			return index
		}
	}

	return -1
}
