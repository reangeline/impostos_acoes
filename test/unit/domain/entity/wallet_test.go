package entity_test

import (
	"testing"

	"github.com/reangeline/impostos_acoes/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestWallet_ProcessBooks(t *testing.T) {
	t.Run("when to create new instance of entity should be 0", func(t *testing.T) {
		wallet := entity.NewWallet()

		assert.Equal(t, 0, wallet.CurrentQuantity)
		assert.Equal(t, 0.0, wallet.WeightedAverage)
	})

	t.Run("when to process buy should be return sum of all my purchases", func(t *testing.T) {
		wallet := entity.NewWallet()

		// w.CurrentQuantity += quantity

		wallet.ProcessBuy(10, 10.0)
		assert.Equal(t, 10, wallet.CurrentQuantity)
		wallet.ProcessBuy(10, 10.0)
		assert.Equal(t, 20, wallet.CurrentQuantity)

	})

	t.Run("when to process buy should be return weighted average of my purchases", func(t *testing.T) {
		wallet := entity.NewWallet()

		wallet.ProcessBuy(10, 10.0)

		assert.Equal(t, 10.0, wallet.WeightedAverage)

		wallet.ProcessBuy(20, 15.0)

		assert.Equal(t, 13.333333333333334, wallet.WeightedAverage)

	})

	t.Run("when to process sell should be return value minus purchase", func(t *testing.T) {
		wallet := entity.NewWallet()

		wallet.ProcessBuy(10, 10.0)

		wallet.ProcessSell(5, 10.0)

		assert.Equal(t, 5, wallet.CurrentQuantity)

	})

	t.Run("when to process sell ", func(t *testing.T) {
		wallet := entity.NewWallet()

		wallet.ProcessBuy(10, 10.0)

		grossProfit, _ := wallet.ProcessSell(5, 20.0)

		// (20 - 10) * float64(5)

		assert.Equal(t, 50.0, grossProfit)

	})

}

// func TestWallet_ProcessBooks(t *testing.T) {

// 	t.Run("should process case 1", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  100,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  15.00,
// 				Quantity:  50,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  15.00,
// 				Quantity:  50,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 0.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 2", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  5.00,
// 				Quantity:  5000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 10000.00, 0.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 3", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  5.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  3000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 1000.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 4", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "buy",
// 				UnitCost:  25.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  15.00,
// 				Quantity:  10000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 0.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 5", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "buy",
// 				UnitCost:  25.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  15.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  25.00,
// 				Quantity:  5000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 0.00, 10000.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 6", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  2.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  2000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  2000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  25.00,
// 				Quantity:  1000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 0.00, 0.00, 3000.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 7", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  2.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  2000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  20.00,
// 				Quantity:  2000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  25.00,
// 				Quantity:  1000,
// 			},
// 			{
// 				Operation: "buy",
// 				UnitCost:  20.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  15.00,
// 				Quantity:  5000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  30.00,
// 				Quantity:  4350,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  30.00,
// 				Quantity:  650,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 0.00, 0.00, 0.00, 3000.00, 0.00, 0.00, 3700.00, 0.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// 	t.Run("should process case 8", func(t *testing.T) {
// 		wallet := entity.Wallet{}
// 		book := []*entity.Book{
// 			{
// 				Operation: "buy",
// 				UnitCost:  10.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  50.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "buy",
// 				UnitCost:  20.00,
// 				Quantity:  10000,
// 			},
// 			{
// 				Operation: "sell",
// 				UnitCost:  50.00,
// 				Quantity:  10000,
// 			},
// 		}

// 		taxes, err := wallet.ProcessBooks(book)

// 		if err != nil {
// 			t.Errorf("Error: %v", err)
// 		}

// 		expectedTaxes := []float64{0.00, 80000.00, 0.00, 60000.00}

// 		assert.Equal(t, expectedTaxes, taxes)

// 	})

// }
