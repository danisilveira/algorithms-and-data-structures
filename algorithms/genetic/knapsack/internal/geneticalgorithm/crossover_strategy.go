package geneticalgorithm

import (
	"math/rand"
	"time"
)

type CrossoverStrategy interface {
	Crossover(Genes, Genes) Genes
}

type InterleavedCrossoverStrategy struct{}

func NewInterleavedCrossoverStrategy() CrossoverStrategy {
	return InterleavedCrossoverStrategy{}
}

func (i InterleavedCrossoverStrategy) Crossover(first Genes, second Genes) Genes {
	genes := make(Genes, len(first))

	for i := 0; i < len(first); i++ {
		gene := first[i]

		if i%2 == 0 {
			gene = second[i]
		}

		genes[i] = gene
	}

	return genes
}

type UniformCrossoverStrategy struct {
	random *rand.Rand
}

func NewUniformCrossoverStrategy() CrossoverStrategy {
	return UniformCrossoverStrategy{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (u UniformCrossoverStrategy) Crossover(first Genes, second Genes) Genes {
	genes := make(Genes, len(first))

	for i := 0; i < len(first); i++ {
		if u.random.Float64() < 0.5 {
			genes[i] = first[i]
			continue
		}

		genes[i] = second[i]
	}

	return genes
}
