package geneticalgorithm

type Configuration struct {
	PopulationSize int
	MaxGenerations int
	CrossoverRate  float64
	MutationRate   float64

	IndividualsComparer     IndividualsComparer
	ParentSelectionStrategy ParentSelectionStrategy
	CrossoverStrategy       CrossoverStrategy
	MutationStrategy        MutationStrategy
}

type OptionFunc func(*Configuration)

func WithPopulationSize(populationSize int) OptionFunc {
	return func(config *Configuration) {
		config.PopulationSize = populationSize
	}
}

func WithMaxGenerations(maxGenerations int) OptionFunc {
	return func(config *Configuration) {
		config.MaxGenerations = maxGenerations
	}
}

func WithCrossoverRate(crossoverRate float64) OptionFunc {
	return func(config *Configuration) {
		config.CrossoverRate = crossoverRate
	}
}

func WithMutationRate(mutationRate float64) OptionFunc {
	return func(config *Configuration) {
		config.MutationRate = mutationRate
	}
}

func WithIndividualsComparer(comparer IndividualsComparer) OptionFunc {
	return func(config *Configuration) {
		config.IndividualsComparer = comparer
	}
}

func WithParentSelectionStrategy(strategy ParentSelectionStrategy) OptionFunc {
	return func(config *Configuration) {
		config.ParentSelectionStrategy = strategy
	}
}

func WithCrossoverStrategy(strategy CrossoverStrategy) OptionFunc {
	return func(config *Configuration) {
		config.CrossoverStrategy = strategy
	}
}

func WithMutationStrategy(strategy MutationStrategy) OptionFunc {
	return func(config *Configuration) {
		config.MutationStrategy = strategy
	}
}
