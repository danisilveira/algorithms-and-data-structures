package geneticalgorithm

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/genetic/knapsack/internal"
)

type Genes []int // Using a bitmap would be better

type Individual struct {
	Genes Genes

	MaxWeight   float64
	TotalValue  float64
	TotalWeight float64
	ItemsCount  int
	Items       []internal.Item

	selectedIndexes []int
}

func NewIndividual(genes Genes, maxWeight float64, items []internal.Item) Individual {
	individual := Individual{
		Genes:           genes,
		MaxWeight:       maxWeight,
		Items:           items,
		selectedIndexes: make([]int, 0, len(items)),
	}

	individual.calculateTotalValueAndTotalWeightAndItemsCount()

	return individual
}

func NewRandomIndividual(maxWeight float64, items []internal.Item) Individual {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	genes := make(Genes, len(items))
	for i := 0; i < len(items); i++ {
		genes[i] = random.Intn(2)
	}

	individual := Individual{
		Genes:           genes,
		MaxWeight:       maxWeight,
		Items:           items,
		selectedIndexes: make([]int, 0, len(items)),
	}

	individual.calculateTotalValueAndTotalWeightAndItemsCount()

	return individual
}

func (i Individual) String() string {
	return fmt.Sprintf("Total Value: %.2f > Total Weight: %.2f > Items Count: %d", i.TotalValue, i.TotalWeight, i.ItemsCount)
}

func (i *Individual) calculateTotalValueAndTotalWeightAndItemsCount() {
	for index, gene := range i.Genes {
		if gene != 1 {
			continue
		}

		item := i.Items[index]
		if (i.TotalWeight + item.Weight) <= i.MaxWeight {
			i.TotalValue += item.Value
			i.TotalWeight += item.Weight
			i.ItemsCount++

			i.selectedIndexes = append(i.selectedIndexes, index)
		}
	}
}
