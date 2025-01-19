package usecase

import (
	"fmt"

	"github.com/reangeline/impostos_acoes/internal/domain/entity"
	"github.com/reangeline/impostos_acoes/internal/dto"
	"github.com/reangeline/impostos_acoes/internal/util"
)

type CalculateTaxesUseCase struct {
	wallet        *entity.Wallet
	globalResults float64
}

func NewCalculateTaxesUseCase() *CalculateTaxesUseCase {
	return &CalculateTaxesUseCase{
		wallet:        entity.NewWallet(),
		globalResults: 0.00,
	}
}

func (c *CalculateTaxesUseCase) Execute(input []*dto.BooksInputDto) ([]*dto.TaxesOutputDto, error) {
	var taxesOutput []*dto.TaxesOutputDto

	for _, book := range input {
		var tax float64
		var err error

		switch book.Operation {
		case "buy":
			c.wallet.ProcessBuy(book.Quantity, book.UnitCost)
			tax = 0.00
		case "sell":
			tax, err = c.logicSell(book.Quantity, book.UnitCost)
			if err != nil {
				return nil, fmt.Errorf("erro ao processar venda: %w", err)
			}

		default:
			return nil, nil
		}

		taxesOutput = append(taxesOutput, &dto.TaxesOutputDto{
			Tax: util.RoundToTwoDecimals(tax),
		})

	}

	return taxesOutput, nil
}

func (c *CalculateTaxesUseCase) logicSell(quantity int, unitCost float64) (float64, error) {

	profit, err := c.wallet.ProcessSell(quantity, unitCost)
	if err != nil {
		return 0.00, err
	}

	netProfit := c.adjustProfitWithGlobalResults(profit)

	valueSell := float64(quantity) * unitCost
	if valueSell > 20000 {
		return c.calculateTax(netProfit), nil
	}

	return 0.00, nil
}

func (c *CalculateTaxesUseCase) adjustProfitWithGlobalResults(profit float64) float64 {
	if profit < 0 {
		c.globalResults += -profit
		return 0.00
	}

	if c.globalResults > 0 {
		remainingProfit := profit - c.globalResults

		if remainingProfit > 0 {
			profit = remainingProfit
			c.globalResults = 0.00
		} else {
			c.globalResults -= profit
			profit = 0.00
		}
	}

	return profit
}

func (c *CalculateTaxesUseCase) calculateTax(netProfit float64) float64 {
	return 0.2 * netProfit
}
