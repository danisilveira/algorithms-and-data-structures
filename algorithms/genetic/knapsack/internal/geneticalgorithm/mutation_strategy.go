package geneticalgorithm

import (
	"math/rand"
	"time"
)

type MutationStrategy interface {
	Mutate(Genes, float64) Genes
}

type SimpleMutationStrategy struct {
	random *rand.Rand
}

func NewSimpleMutationStrategy() MutationStrategy {
	return SimpleMutationStrategy{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

}

func (s SimpleMutationStrategy) Mutate(genes Genes, mutationRate float64) Genes {
	newGenes := make(Genes, len(genes))

	for i := 0; i < len(genes); i++ {
		if s.random.Float64() >= mutationRate {
			continue
		}

		newGene := 1
		if genes[i] == 1 {
			newGene = 0
		}

		newGenes[i] = newGene
	}

	return newGenes
}
