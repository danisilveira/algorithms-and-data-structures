package insertionsort

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](items []T) {
	for i := 1; i < len(items); i++ {
		tempValue := items[i]
		position := i - 1

		for position >= 0 {
			if items[position] <= tempValue {
				break
			}

			items[position+1] = items[position]
			position--
		}

		items[position+1] = tempValue
	}
}
