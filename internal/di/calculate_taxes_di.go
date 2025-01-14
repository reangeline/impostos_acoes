package di

import (
	"github.com/reangeline/impostos_acoes/internal/domain/usecase"
	"github.com/reangeline/impostos_acoes/internal/presentation/handler"
)

func CalculateTaxesInjection() *handler.StdinHandler {

	calculateTaxesUsecase := usecase.NewCalculateTaxesUseCase()

	calculate := handler.NewCalculateTaxesHandler(calculateTaxesUsecase)

	return calculate

}
