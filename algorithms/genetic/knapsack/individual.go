package knapsack

import (
	"fmt"
	"math/rand"
	"time"
)

type individual struct {
	knapsack Knapsack
	items    []Item

	itemsCount  int
	totalWeight float64
	totalValue  float64

	genes []int // using a bitset would be better
}

func newRandomIndividual(config config) individual {
	rand.Seed(time.Now().UnixNano())

	genes := make([]int, len(config.items))
	for i := 0; i < len(config.items); i++ {
		genes[i] = rand.Intn(2)
	}

	individual := individual{
		knapsack: *config.knapsack,
		items:    config.items,

		genes: genes,
	}

	individual.compute()

	return individual
}

func (i individual) String() string {
	return fmt.Sprintf("Individual has %d items. Total weight of %.2f. Total value of %.2f.", i.itemsCount, i.totalWeight, i.totalValue)
}

func (i *individual) compute() {
	for index, gene := range i.genes {
		if gene == 0 {
			continue
		}

		item := i.items[index]

		if (i.totalWeight + item.Weight) <= i.knapsack.MaxWeight {
			i.totalWeight += item.Weight
			i.totalValue += item.Value
			i.itemsCount++
		}
	}
}
