package entity

type Wallet struct {
	CurrentQuantity int
	WeightedAverage float64
	Results         float64
}

func (w *Wallet) ProcessBooks(books []*Book) ([]float64, error) {
	resultTaxes := make([]float64, len(books))

	for index, book := range books {
		tax, err := w.ProcessBook(book)
		if err != nil {
			return nil, err
		}
		resultTaxes[index] = tax
	}

	return resultTaxes, nil
}

func (w *Wallet) ProcessBook(book *Book) (float64, error) {
	switch book.Operation {
	case "buy":
		w.ProcessBuy(book.Quantity, book.UnitCost)
		return 0.00, nil
	case "sell":
		return w.ProcessSell(book.Quantity, book.UnitCost)
	default:
		return 0.00, nil
	}
}

func (w *Wallet) ProcessBuy(quantity int, unitCost float64) {

	newWeightedAverage := ((float64(w.CurrentQuantity) * w.WeightedAverage) + (float64(quantity) * unitCost)) / float64(w.CurrentQuantity+quantity)
	w.CurrentQuantity += quantity
	w.WeightedAverage = newWeightedAverage
}

func (w *Wallet) ProcessSell(quantity int, unitCost float64) (float64, error) {
	valueSell := float64(quantity) * unitCost

	w.CurrentQuantity -= quantity
	if unitCost < w.WeightedAverage {
		loss := (w.WeightedAverage - unitCost) * float64(quantity)
		w.Results += loss
		return 0.00, nil
	}

	grossProfit := (unitCost - w.WeightedAverage) * float64(quantity)

	netProfit := grossProfit
	if w.Results > 0.00 {
		if grossProfit > w.Results {
			netProfit -= w.Results
			w.Results = 0.00
		} else {
			w.Results -= grossProfit
			netProfit = 0.00
		}
	}

	if valueSell <= 20000 {
		return 0.00, nil
	}

	tax := 0.2 * netProfit

	return tax, nil

}
