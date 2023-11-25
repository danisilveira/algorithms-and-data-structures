package mergesort

import "cmp"

func Sort[T cmp.Ordered](items []T) []T {
	if len(items) == 1 {
		return items
	}

	firstHalf := Sort(items[:len(items)/2])
	secondHalf := Sort(items[len(items)/2:])

	return merge(firstHalf, secondHalf)
}

func merge[T cmp.Ordered](firstHalf, secondHalf []T) []T {
	items := []T{}

	firstHalfIndex := 0
	secondHalfIndex := 0

	for firstHalfIndex < len(firstHalf) && secondHalfIndex < len(secondHalf) {
		if firstHalf[firstHalfIndex] < secondHalf[secondHalfIndex] {
			items = append(items, firstHalf[firstHalfIndex])
			firstHalfIndex++
		} else {
			items = append(items, secondHalf[secondHalfIndex])
			secondHalfIndex++
		}
	}

	for firstHalfIndex < len(firstHalf) {
		items = append(items, firstHalf[firstHalfIndex])
		firstHalfIndex++
	}

	for secondHalfIndex < len(secondHalf) {
		items = append(items, secondHalf[secondHalfIndex])
		secondHalfIndex++
	}

	return items
}
