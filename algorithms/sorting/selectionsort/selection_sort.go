package selectionsort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](items []T) {
	for i := 0; i < len(items)-1; i++ {
		lowest := i

		for j := i + 1; j < len(items); j++ {
			if items[j] < items[lowest] {
				lowest = j
			}
		}

		if lowest != i {
			items[i], items[lowest] = items[lowest], items[i]
		}
	}
}
