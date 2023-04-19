package knapsack

import (
	"github.com/danisilveira/algorithms-and-data-structures/algorithms/genetic/knapsack/internal"
	"github.com/danisilveira/algorithms-and-data-structures/algorithms/genetic/knapsack/internal/geneticalgorithm"
)

type Knapsack struct {
	MaxWeight float64

	Items []Item
}

func New(maxWeight float64) Knapsack {
	return Knapsack{
		MaxWeight: maxWeight,
	}
}

func (k *Knapsack) SelectAndPutInside(items []Item) {
	internalItems := make([]internal.Item, len(items))
	for i, item := range items {
		internalItems[i] = internal.Item{
			Value:  item.Value,
			Weight: item.Weight,
		}
	}

	geneticAlgorithm := geneticalgorithm.NewWithOptions(
		geneticalgorithm.WithPopulationSize(200),
		geneticalgorithm.WithMaxGenerations(1000),
	)
	indexes := geneticAlgorithm.Run(k.MaxWeight, internalItems)

	itemsToPutInside := make([]Item, 0, len(indexes))
	for _, index := range indexes {
		itemsToPutInside = append(itemsToPutInside, items[index])
	}

	k.Items = itemsToPutInside
}

func (k *Knapsack) TotalValue() float64 {
	totalValue := 0.0
	for _, item := range k.Items {
		totalValue += item.Value
	}

	return totalValue
}

func (k *Knapsack) TotalWeight() float64 {
	totalWeight := 0.0
	for _, item := range k.Items {
		totalWeight += item.Weight
	}

	return totalWeight
}
