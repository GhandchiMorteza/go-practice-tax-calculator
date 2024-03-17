package main

import (
	"fmt"

	"example.com/price-calculator/data"
	"example.com/price-calculator/storage"
	"example.com/price-calculator/tax"
)

func main() {
	filePath := "data/prices.txt"
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	prices, err := data.LoadPrices(filePath)
	if err != nil {
		fmt.Printf("Failed to load data: %v\n", err)
		return
	}

	for _, rate := range taxRates {
		job := tax.NewTaxIncludedPriceJob(rate, prices)
		job.Process()

		fileName := fmt.Sprintf("result_%.0f.json", rate*100)
		if err := storage.SaveToJSON(job.TaxIncludedPrices, fileName); err != nil {
			fmt.Printf("Error saving results to JSON: %v\n", err)
			return
		}
	}
}