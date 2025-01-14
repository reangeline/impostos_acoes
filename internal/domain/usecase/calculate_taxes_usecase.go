package usecase

import (
	"github.com/reangeline/impostos_acoes/internal/domain/entity"
	"github.com/reangeline/impostos_acoes/internal/dto"
	"github.com/reangeline/impostos_acoes/internal/util"
)

type CalculateTaxesUseCase struct {
}

func NewCalculateTaxesUseCase() *CalculateTaxesUseCase {
	return &CalculateTaxesUseCase{}
}

func (c *CalculateTaxesUseCase) Execute(input []*dto.BooksInputDto) ([]*dto.TaxesOutputDto, error) {
	var books []*entity.Book
	var wallet *entity.Wallet

	for _, book := range input {
		bookEntity := entity.Book{
			Operation: book.Operation,
			UnitCost:  book.UnitCost,
			Quantity:  book.Quantity,
		}

		books = append(books, &bookEntity)
	}

	wallet = &entity.Wallet{}
	taxes, err := wallet.ProcessBooks(books)
	if err != nil {
		return nil, err
	}

	var taxesOutput []*dto.TaxesOutputDto
	for _, tax := range taxes {
		roundTax := util.RoundToTwoDecimals(tax)
		taxOutput := dto.TaxesOutputDto{
			Tax: roundTax,
		}

		taxesOutput = append(taxesOutput, &taxOutput)
	}

	return taxesOutput, nil

}
