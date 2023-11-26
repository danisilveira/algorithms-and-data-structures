package binarysearch_test

import (
	"fmt"
	"testing"

	"github.com/danisilveira/algorithms-and-data-structures/datastructures/tree/binarysearch"
)

func TestBinarySearchTree(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tr := binarysearch.NewTree[int]()
		tr.AddValue(50)
		tr.AddValue(25)
		tr.AddValue(75)
		tr.AddValue(20)
		tr.AddValue(30)
		tr.AddValue(70)
		tr.AddValue(80)
		tr.AddValue(27)

		it := tr.InOrder()
		for it.HasNext() {
			fmt.Println(it.Value())
		}

		fmt.Println()

		it = tr.PreOrder()
		for it.HasNext() {
			fmt.Println(it.Value())
		}

		fmt.Println()

		it = tr.PostOrder()
		for it.HasNext() {
			fmt.Println(it.Value())
		}

		t.Fail()
	})
}
