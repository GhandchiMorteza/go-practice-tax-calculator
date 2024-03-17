package main

import (
	"fmt"

	"example.com/price-calculator/data"
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

	for _, taxRate := range taxRates {
		priceJob := tax.NewTaxIncludedPriceJob(taxRate, prices)
		priceJob.Process()
	}

}