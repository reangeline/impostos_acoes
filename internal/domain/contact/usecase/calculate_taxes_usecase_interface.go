package contract

import "github.com/reangeline/impostos_acoes/internal/dto"

type CalculateTaxesUseCaseInterface interface {
	Execute(input []*dto.BooksInputDto) ([]*dto.TaxesOutputDto, error)
}
