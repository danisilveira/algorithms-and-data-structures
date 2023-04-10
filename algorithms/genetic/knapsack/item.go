package knapsack

type Item struct {
	Name   string
	Weight float64
	Value  float64
}

func NewItem(name string, weight, value float64) Item {
	return Item{
		Name:   name,
		Weight: weight,
		Value:  value,
	}
}
