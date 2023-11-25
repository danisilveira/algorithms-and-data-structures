package inversions

import "cmp"

func Count[S ~[]T, T cmp.Ordered](items S) int {
	_, count := sortAndCountInversions(items)
	return count
}

func sortAndCountInversions[S ~[]T, T cmp.Ordered](items S) (orderedItems S, inversionsCount int) {
	if len(items) <= 1 {
		return items, 0
	}

	firstHalf, leftInversions := sortAndCountInversions(items[:len(items)/2])
	secondHalf, rightInversions := sortAndCountInversions(items[len(items)/2:])
	merged, splitInversions := mergeAndCountSplitInversions(firstHalf, secondHalf)

	return merged, (leftInversions + rightInversions + splitInversions)
}

func mergeAndCountSplitInversions[S ~[]T, T cmp.Ordered](firstHalf, secondHalf S) (orderedItems S, inversionsCount int) {
	items := []T{}
	splitInversions := 0

	firstHalfIndex := 0
	secondHalfIndex := 0

	for firstHalfIndex < len(firstHalf) && secondHalfIndex < len(secondHalf) {
		if firstHalf[firstHalfIndex] < secondHalf[secondHalfIndex] {
			items = append(items, firstHalf[firstHalfIndex])
			firstHalfIndex++
		} else {
			items = append(items, secondHalf[secondHalfIndex])
			secondHalfIndex++
			splitInversions += (len(firstHalf) - firstHalfIndex)
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

	return items, splitInversions
}
