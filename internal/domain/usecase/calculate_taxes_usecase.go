package usecase

import (
	"github.com/reangeline/impostos_acoes/internal/domain/entity"
	"github.com/reangeline/impostos_acoes/internal/dto"
	"github.com/reangeline/impostos_acoes/internal/util"
)

type CalculateTaxesUseCase struct {
	globalResults float64
	wallet        *entity.Wallet
}

func NewCalculateTaxesUseCase() *CalculateTaxesUseCase {
	return &CalculateTaxesUseCase{
		globalResults: 0,
		wallet:        entity.NewWallet(),
	}
}

func (c *CalculateTaxesUseCase) Execute(input []*dto.BooksInputDto) ([]*dto.TaxesOutputDto, error) {
	var taxesOutput []*dto.TaxesOutputDto

	for _, book := range input {
		var tax float64
		switch book.Operation {
		case "buy":
			c.wallet.ProcessBuy(book.Quantity, book.UnitCost)
			tax = 0
		case "sell":
			tax = c.logicSell(book.Quantity, book.UnitCost)
		}

		roundTax := util.RoundToTwoDecimals(tax)
		taxOutput := dto.TaxesOutputDto{
			Tax: roundTax,
		}

		taxesOutput = append(taxesOutput, &taxOutput)
	}

	return taxesOutput, nil

}

func (c *CalculateTaxesUseCase) logicSell(quantity int, unitCost float64) float64 {
	grossProfit, _ := c.wallet.ProcessSell(quantity, unitCost)

	if unitCost < c.wallet.WeightedAverage {
		loss := (c.wallet.WeightedAverage - unitCost) * float64(quantity)
		c.globalResults += loss
		return 0
	}

	netProfit := c.adjustGlobalResults(grossProfit)

	valueSell := float64(quantity) * unitCost
	if valueSell <= 20000 {
		return 0.00
	}

	tax := 0.2 * netProfit
	return tax

}

func (c *CalculateTaxesUseCase) adjustGlobalResults(grossProfit float64) float64 {

	netProfit := grossProfit
	if c.globalResults > 0.00 {
		if grossProfit > c.globalResults {
			netProfit -= c.globalResults
			c.globalResults = 0.00
		} else {
			c.globalResults -= grossProfit
			netProfit = 0.00
		}
	}

	return netProfit

}
