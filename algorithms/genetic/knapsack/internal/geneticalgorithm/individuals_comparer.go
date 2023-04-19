package geneticalgorithm

type IndividualsComparer interface {
	Compare(Individual, Individual) bool
}

type ValueAndWeightAndItemsCount struct{}

func NewValueAndWeightAndItemsCountComparer() IndividualsComparer {
	return ValueAndWeightAndItemsCount{}
}

func (v ValueAndWeightAndItemsCount) Compare(first Individual, second Individual) bool {
	if first.TotalValue == second.TotalValue {
		if first.TotalWeight == second.TotalWeight {
			return first.ItemsCount >= second.ItemsCount
		}

		return first.TotalWeight <= second.TotalWeight
	}

	return first.TotalValue >= second.TotalValue
}
