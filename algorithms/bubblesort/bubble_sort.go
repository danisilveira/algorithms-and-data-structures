package bubblesort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](items []T) {
	sorted := false

	for !sorted {
		sorted = true
		unsortedUntilIndex := len(items) - 1

		for j := 0; j < unsortedUntilIndex; j++ {
			if items[j] > items[j+1] {
				temp := items[j]
				items[j] = items[j+1]
				items[j+1] = temp

				sorted = false
			}
		}

		unsortedUntilIndex--
	}
}
