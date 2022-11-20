package quickselect

import "golang.org/x/exp/constraints"

func Select[T constraints.Ordered](items []T, kthLowestValue int) T {
	return qselect(items, kthLowestValue-1, 0, len(items)-1)
}

func qselect[T constraints.Ordered](items []T, kthLowestValue, leftPointer, rightPointer int) T {
	if rightPointer-leftPointer <= 0 {
		return items[leftPointer]
	}

	pivotIndex := partition(items, leftPointer, rightPointer)

	if kthLowestValue < pivotIndex {
		return qselect(items, kthLowestValue, leftPointer, pivotIndex-1)
	}

	if kthLowestValue > pivotIndex {
		return qselect(items, kthLowestValue, pivotIndex+1, rightPointer)
	}

	return items[pivotIndex]
}

func partition[T constraints.Ordered](items []T, leftPointer, rightPointer int) int {
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
