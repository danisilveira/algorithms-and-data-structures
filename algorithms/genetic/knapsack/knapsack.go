package knapsack

import "fmt"

type Knapsack struct {
	MaxWeight float64
}

func New(maxWeight float64) Knapsack {
	return Knapsack{
		MaxWeight: maxWeight,
	}
}

func (k *Knapsack) SelectAndPutInside(items []Item) {
	config := config{
		knapsack: k,
		items:    items,

		populationSize:  20,
		maxGenerations:  100,
		fitnessComparer: newValueAndWeightAndItemsCountComparer(),
	}

	population := newPopulationWithRandomIndividuals(config)
	individual := population.GetBestIndividual()
	fmt.Printf("The best individual is: %s\n\n", individual)

	population.Print()
	// for i := 0; i < config.maxGenerations; i++ {
	// 	population = population.Evolve()
	// }
}
