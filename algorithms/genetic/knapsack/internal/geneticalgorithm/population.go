package geneticalgorithm

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/danisilveira/algorithms-and-data-structures/algorithms/genetic/knapsack/internal"
)

type Population struct {
	Individuals []Individual

	MaxWeight float64
	Items     []internal.Item
	Config    *Configuration

	random *rand.Rand
}

func NewPopulation(individuals []Individual, maxWeight float64, items []internal.Item, config *Configuration) Population {
	population := Population{
		Individuals: individuals,

		MaxWeight: maxWeight,
		Items:     items,
		Config:    config,

		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	population.sortIndividuals()

	return population
}

func NewRandomPopulation(maxWeight float64, items []internal.Item, config *Configuration) Population {
	individuals := make([]Individual, config.PopulationSize)
	for i := 0; i < config.PopulationSize; i++ {
		individuals[i] = NewRandomIndividual(maxWeight, items)
	}

	population := Population{
		Individuals: individuals,

		MaxWeight: maxWeight,
		Items:     items,
		Config:    config,

		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	population.sortIndividuals()

	return population
}

func (p Population) Evolve() Population {
	newIndividuals := make([]Individual, 0, p.Config.PopulationSize)
	newIndividuals = append(newIndividuals, p.BestIndividuals(p.Config.PopulationSize/10)...)

	for i := len(newIndividuals); i < p.Config.PopulationSize; i++ {
		parent1 := p.Config.ParentSelectionStrategy.SelectParent(p)
		parent2 := p.Config.ParentSelectionStrategy.SelectParent(p)

		var genes Genes
		if p.random.Float64() < p.Config.CrossoverRate {
			genes = p.Config.CrossoverStrategy.Crossover(parent1.Genes, parent2.Genes)
		} else {
			genes = parent1.Genes
		}

		genes = p.Config.MutationStrategy.Mutate(genes, p.Config.MutationRate)
		child := NewIndividual(genes, p.MaxWeight, p.Items)
		newIndividuals = append(newIndividuals, child)
	}

	return NewPopulation(newIndividuals, p.MaxWeight, p.Items, p.Config)
}

func (p Population) BestIndividuals(count int) []Individual {
	return p.Individuals[:count]
}

func (p Population) BestIndividual() Individual {
	return p.Individuals[0]
}

func (p Population) Print() {
	for _, individual := range p.Individuals {
		fmt.Println(individual)
	}
}

func (p Population) sortIndividuals() {
	sort.SliceStable(p.Individuals, func(i, j int) bool {
		first := p.Individuals[i]
		second := p.Individuals[j]

		return p.Config.IndividualsComparer.Compare(first, second)
	})
}
