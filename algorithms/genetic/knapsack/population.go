package knapsack

import (
	"fmt"
	"sort"
)

type population struct {
	individuals []individual

	config config
}

func newPopulation(individuals []individual, config config) population {
	population := population{
		individuals: individuals,
		config:      config,
	}

	population.sortIndividuals()

	return population
}

func newPopulationWithRandomIndividuals(config config) population {
	individuals := make([]individual, config.populationSize)
	for i := 0; i < config.populationSize; i++ {
		individuals[i] = newRandomIndividual(config)
	}

	population := population{
		individuals: individuals,
		config:      config,
	}

	population.sortIndividuals()

	return population
}

func (p population) Evolve() population {
	newIndividuals := make([]individual, 0, len(p.individuals))
	newIndividuals = append(newIndividuals, p.GetBestIndividuals(p.config.populationSize/10)...)

	for len(newIndividuals) < p.config.populationSize {
		// select parent 1
		// select parent 2
		// crossover
		// mutate
		// add individual to newIndividuals slice
	}

	return newPopulation(newIndividuals, p.config)
}

func (p population) GetBestIndividual() individual {
	return p.individuals[0]
}

func (p population) GetBestIndividuals(count int) []individual {
	return p.individuals[:count]
}

func (p population) Print() {
	for _, individual := range p.individuals {
		fmt.Println(individual)
	}
}

func (p population) sortIndividuals() {
	sort.SliceStable(p.individuals, func(i, j int) bool {
		firstIndividual := p.individuals[i]
		secondIndividual := p.individuals[j]

		return p.config.fitnessComparer.Compare(firstIndividual, secondIndividual)
	})
}
