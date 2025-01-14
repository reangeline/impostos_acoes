package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	contract "github.com/reangeline/impostos_acoes/internal/domain/contact/usecase"
	"github.com/reangeline/impostos_acoes/internal/dto"
)

type StdinHandler struct {
	calculateTaxesUseCase contract.CalculateTaxesUseCaseInterface
}

func NewCalculateTaxesHandler(calculateTaxesUseCase contract.CalculateTaxesUseCaseInterface) *StdinHandler {
	return &StdinHandler{
		calculateTaxesUseCase: calculateTaxesUseCase,
	}
}

func (ct *StdinHandler) HandleInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()

		var books []*dto.BooksInputDto
		if err := json.Unmarshal([]byte(row), &books); err != nil {
			fmt.Fprintln(os.Stderr, "Error decoding JSON input:", err)
			continue
		}

		resultTaxes, err := ct.calculateTaxesUseCase.Execute(books)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error processing operations:", err)
			continue
		}

		output, err := json.Marshal(resultTaxes)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error encoding JSON output:", err)
			continue
		}

		os.Stdout.WriteString(string(output) + "\n")

	}

}
