package entity_test

import (
	"fmt"
	"testing"

	"github.com/reangeline/impostos_acoes/internal/domain/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	wallet := entity.NewWallet()

	assert.NotNil(t, wallet)
	assert.Equal(t, 0, wallet.CurrentQuantity)
	assert.Equal(t, 0.00, wallet.WeightedAverage)
	assert.Equal(t, 0.00, wallet.Results)
}

func TestProcessBuy(t *testing.T) {
	wallet := entity.NewWallet()

	wallet.ProcessBuy(10, 100.00)
	assert.Equal(t, 10, wallet.CurrentQuantity)
	assert.Equal(t, 100.00, wallet.WeightedAverage)

	wallet.ProcessBuy(5, 200.00)
	assert.Equal(t, 15, wallet.CurrentQuantity)
	assert.Equal(t, 133.33333333333334, wallet.WeightedAverage)
}

func TestProcessSell(t *testing.T) {
	wallet := entity.NewWallet()
	wallet.ProcessBuy(10, 100.00)

	fmt.Println(wallet.WeightedAverage)

	grossProfit, err := wallet.ProcessSell(5, 150.00)
	assert.Nil(t, err)

	assert.Equal(t, 250.00, grossProfit)

	assert.Equal(t, 5, wallet.CurrentQuantity)
	assert.Equal(t, 100.00, wallet.WeightedAverage)
}
