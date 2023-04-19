package geneticalgorithm

import (
	"fmt"
	"math"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/genetic/knapsack/internal"
)

type GeneticAlgorithm struct {
	Config *Configuration
}

func New() GeneticAlgorithm {
	individualsComparer := NewValueAndWeightAndItemsCountComparer()
	parentSelectionStrategy := NewTournamentParentSelectionStrategy(10, individualsComparer)
	crossoverStrategy := NewUniformCrossoverStrategy()
	mutationStrategy := NewSimpleMutationStrategy()

	return GeneticAlgorithm{
		Config: &Configuration{
			PopulationSize:          100,
			MaxGenerations:          500,
			CrossoverRate:           0.7,
			MutationRate:            0.1,
			IndividualsComparer:     individualsComparer,
			ParentSelectionStrategy: parentSelectionStrategy,
			CrossoverStrategy:       crossoverStrategy,
			MutationStrategy:        mutationStrategy,
		},
	}
}

func NewWithOptions(configurationOptions ...OptionFunc) GeneticAlgorithm {
	geneticAlgorithm := New()

	for _, option := range configurationOptions {
		option(geneticAlgorithm.Config)
	}

	return geneticAlgorithm
}

func NewWithConfig(config Configuration) GeneticAlgorithm {
	return GeneticAlgorithm{
		Config: &config,
	}
}

func (ga *GeneticAlgorithm) Run(maxWeight float64, items []internal.Item) []int {
	population := NewRandomPopulation(maxWeight, items, ga.Config)
	bestIndividual := population.BestIndividual()

	fmt.Printf("Config:\n%v\n\n", *ga.Config)
	fmt.Printf("First Generation: %v\n", bestIndividual)

	initialMutationRate := ga.Config.MutationRate
	staleGenerationsCount := 0

	for i := 0; i < ga.Config.MaxGenerations; i++ {
		population = population.Evolve()
		candidate := population.BestIndividual()

		isIndividualBetterThanTheCandidate := ga.Config.IndividualsComparer.Compare(bestIndividual, candidate)
		if isIndividualBetterThanTheCandidate {
			staleGenerationsCount++
			ga.Config.MutationRate = initialMutationRate * (math.Pow(1.1, float64(staleGenerationsCount)))
		} else {
			bestIndividual = candidate
			ga.Config.MutationRate = initialMutationRate
			staleGenerationsCount = 0
		}
	}

	fmt.Printf("Solution: %v\n", bestIndividual)

	return bestIndividual.selectedIndexes
}
