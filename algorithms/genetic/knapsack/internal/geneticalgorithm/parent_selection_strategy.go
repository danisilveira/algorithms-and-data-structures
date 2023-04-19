package geneticalgorithm

import (
	"math/rand"
	"sort"
	"time"
)

type ParentSelectionStrategy interface {
	SelectParent(Population) Individual
}

type TournamentParentSelectionStrategy struct {
	tournamentSize      int
	individualsComparer IndividualsComparer
	random              *rand.Rand
}

func NewTournamentParentSelectionStrategy(tournamentSize int, individualsComparer IndividualsComparer) ParentSelectionStrategy {
	return TournamentParentSelectionStrategy{
		tournamentSize:      tournamentSize,
		individualsComparer: individualsComparer,
		random:              rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (t TournamentParentSelectionStrategy) SelectParent(population Population) Individual {
	individuals := make([]Individual, t.tournamentSize)

	for i := 0; i < t.tournamentSize; i++ {
		individuals[i] = population.Individuals[t.random.Intn(len(population.Individuals))]
	}

	sort.SliceStable(individuals, func(i, j int) bool {
		first := individuals[i]
		second := individuals[j]

		return t.individualsComparer.Compare(first, second)
	})

	return individuals[0]
}
