package entity

type Wallet struct {
	CurrentQuantity int
	WeightedAverage float64
}

func NewWallet() *Wallet {
	return &Wallet{
		CurrentQuantity: 0,
		WeightedAverage: 0.0,
	}
}

func (w *Wallet) ProcessBuy(quantity int, unitCost float64) {

	newWeightedAverage := ((float64(w.CurrentQuantity) * w.WeightedAverage) + (float64(quantity) * unitCost)) / float64(w.CurrentQuantity+quantity)
	w.CurrentQuantity += quantity
	w.WeightedAverage = newWeightedAverage
}

func (w *Wallet) ProcessSell(quantity int, unitCost float64) (float64, error) {
	w.CurrentQuantity -= quantity

	grossProfit := (unitCost - w.WeightedAverage) * float64(quantity)

	return grossProfit, nil

}
