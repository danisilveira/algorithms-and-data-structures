package knapsack

type config struct {
	knapsack *Knapsack
	items    []Item

	populationSize  int
	maxGenerations  int
	fitnessComparer fitnessComparer
}
