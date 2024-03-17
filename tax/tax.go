package tax

import "fmt"

type TaxIncludedPriceJob struct {
	TaxRate           float64 					`json:"tax_rate"`
	InputPrices       []float64 				`json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(taxRate float64, prices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: prices,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result
}