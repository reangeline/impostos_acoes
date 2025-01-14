package dto

type BooksInputDto struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}
