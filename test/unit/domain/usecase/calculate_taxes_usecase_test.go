package usecase_test

import (
	"testing"

	"github.com/reangeline/impostos_acoes/internal/domain/usecase"
	"github.com/reangeline/impostos_acoes/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestCalculateTaxesUseCase_Execute_TableDrivenTest(t *testing.T) {
	tests := []struct {
		name          string
		input         []*dto.BooksInputDto
		expectedTaxes []float64
		expectedError error
	}{
		{
			name: "should process case 1",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 100},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
				{Operation: "sell", UnitCost: 15.00, Quantity: 50},
			},
			expectedTaxes: []float64{0.00, 0.00, 0.00},
		},
		{
			name: "should process case 2",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			},
			expectedTaxes: []float64{0.00, 10000.00, 0.00},
		},
		{
			name: "should process case 3",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
			},
			expectedTaxes: []float64{0.00, 0.00, 1000.00},
		},
		{
			name: "should process case 4",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			},
			expectedTaxes: []float64{0.00, 0.00, 0.00},
		},
		{
			name: "should process case 5",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
			},
			expectedTaxes: []float64{0.00, 0.00, 0.00, 10000.00},
		},
		{
			name: "should process case 6",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			},
			expectedTaxes: []float64{0.00, 0.00, 0.00, 0.00, 3000.00},
		},
		{
			name: "should process case 7",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
				{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
				{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
				{Operation: "sell", UnitCost: 30.00, Quantity: 650},
			},
			expectedTaxes: []float64{0.00, 0.00, 0.00, 0.00, 3000.00, 0.00, 0.00, 3700.00, 0.00},
		},
		{
			name: "should process case 8",
			input: []*dto.BooksInputDto{
				{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
				{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
				{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			},
			expectedTaxes: []float64{0.00, 80000.00, 0.00, 60000.00},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := usecase.NewCalculateTaxesUseCase()
			taxes, err := useCase.Execute(tt.input)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}

			var resultTaxes []float64
			for _, tax := range taxes {
				resultTaxes = append(resultTaxes, tax.Tax)
			}

			assert.Equal(t, tt.expectedTaxes, resultTaxes)
		})
	}
}

func TestCalculateTaxesUseCase_Execute_AAA(t *testing.T) {
	t.Run("should process case 1", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 100},
			{Operation: "sell", UnitCost: 15.00, Quantity: 50},
			{Operation: "sell", UnitCost: 15.00, Quantity: 50},
		}
		expectedTaxes := []float64{0.00, 0.00, 0.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 2", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
		}
		expectedTaxes := []float64{0.00, 10000.00, 0.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 3", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 5.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 3000},
		}
		expectedTaxes := []float64{0.00, 0.00, 1000.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 4", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
		}
		expectedTaxes := []float64{0.00, 0.00, 0.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 5", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "buy", UnitCost: 25.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 15.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 25.00, Quantity: 5000},
		}
		expectedTaxes := []float64{0.00, 0.00, 0.00, 10000.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 6", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
			{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
		}
		expectedTaxes := []float64{0.00, 0.00, 0.00, 0.00, 3000.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 7", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 2.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
			{Operation: "sell", UnitCost: 20.00, Quantity: 2000},
			{Operation: "sell", UnitCost: 25.00, Quantity: 1000},
			{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 15.00, Quantity: 5000},
			{Operation: "sell", UnitCost: 30.00, Quantity: 4350},
			{Operation: "sell", UnitCost: 30.00, Quantity: 650},
		}
		expectedTaxes := []float64{0.00, 0.00, 0.00, 0.00, 3000.00, 0.00, 0.00, 3700.00, 0.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})

	t.Run("should process case 8", func(t *testing.T) {
		// Arrange
		useCase := usecase.NewCalculateTaxesUseCase()
		input := []*dto.BooksInputDto{
			{Operation: "buy", UnitCost: 10.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
			{Operation: "buy", UnitCost: 20.00, Quantity: 10000},
			{Operation: "sell", UnitCost: 50.00, Quantity: 10000},
		}
		expectedTaxes := []float64{0.00, 80000.00, 0.00, 60000.00}

		// Act
		taxes, err := useCase.Execute(input)

		// Assert
		assert.NoError(t, err)
		var resultTaxes []float64
		for _, tax := range taxes {
			resultTaxes = append(resultTaxes, tax.Tax)
		}
		assert.Equal(t, expectedTaxes, resultTaxes)
	})
}
