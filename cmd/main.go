package main

import "github.com/reangeline/impostos_acoes/internal/di"

func main() {

	handle := di.CalculateTaxesInjection()

	handle.HandleInput()

}
