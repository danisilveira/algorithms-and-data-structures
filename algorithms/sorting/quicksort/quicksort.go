package quicksort

import "cmp"

func Sort[T cmp.Ordered](items []T) {
	sort(items, 0, len(items)-1)
}

func sort[T cmp.Ordered](items []T, leftPointer, rightPointer int) {
	if rightPointer-leftPointer <= 0 {
		return
	}

	pivotPointer := partition(items, leftPointer, rightPointer)

	sort(items, leftPointer, pivotPointer-1)
	sort(items, pivotPointer+1, rightPointer)
}

func partition[T cmp.Ordered](items []T, leftPointer, rightPointer int) int {
	pivot := items[rightPointer]

	for currentPointer := leftPointer; currentPointer < rightPointer; currentPointer++ {
		if items[currentPointer] < pivot {
			items[leftPointer], items[currentPointer] = items[currentPointer], items[leftPointer]
			leftPointer++
		}
	}

	items[leftPointer], items[rightPointer] = items[rightPointer], items[leftPointer]

	return leftPointer
}
