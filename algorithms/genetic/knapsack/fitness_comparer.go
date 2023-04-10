package knapsack

type fitnessComparer interface {
	Compare(firstIndividual, secondIndividual individual) bool
}

type weightComparer struct{}

func newWeightComparer() fitnessComparer {
	return weightComparer{}
}

func (c weightComparer) Compare(firstIndividual, secondIndividual individual) bool {
	return firstIndividual.totalWeight < secondIndividual.totalWeight
}

type valueAndWeightComparer struct{}

func newValueAndWeightComparer() fitnessComparer {
	return valueAndWeightComparer{}
}

func (c valueAndWeightComparer) Compare(firstIndividual, secondIndividual individual) bool {
	if firstIndividual.totalValue == secondIndividual.totalValue {
		return firstIndividual.totalWeight < secondIndividual.totalWeight
	}

	return firstIndividual.totalValue > secondIndividual.totalValue
}

type valueAndWeightAndItemsCountComparer struct{}

func newValueAndWeightAndItemsCountComparer() fitnessComparer {
	return valueAndWeightAndItemsCountComparer{}
}

func (c valueAndWeightAndItemsCountComparer) Compare(firstIndividual, secondIndividual individual) bool {
	if firstIndividual.totalValue == secondIndividual.totalValue {
		if firstIndividual.totalWeight == secondIndividual.totalWeight {
			return firstIndividual.itemsCount > secondIndividual.itemsCount
		}

		return firstIndividual.totalWeight < secondIndividual.totalWeight
	}

	return firstIndividual.totalValue > secondIndividual.totalValue
}
