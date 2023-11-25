package insertionsort

import "cmp"

func Sort[T cmp.Ordered](items []T) {
	for i := 1; i < len(items); i++ {
		tempValue := items[i]
		position := i - 1

		for position >= 0 && items[position] > tempValue {
			items[position+1] = items[position]
			position--
		}

		items[position+1] = tempValue
	}
}
