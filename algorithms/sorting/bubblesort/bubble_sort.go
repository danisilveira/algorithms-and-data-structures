package bubblesort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](items []T) {
	sorted := false

	unsortedUntilIndex := len(items) - 1
	for !sorted {
		sorted = true

		for j := 0; j < unsortedUntilIndex; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
				sorted = false
			}
		}

		unsortedUntilIndex--
	}
}
