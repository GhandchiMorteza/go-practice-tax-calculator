package main

import (
	"fmt"
	"os"

	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob, err := prices.NewTaxIncludedPriceJob(taxRate)
		if err != nil {
			fmt.Printf("Failed to load prices: %v\n", err)
      os.Exit(1)
		}
		priceJob.Process()
	}

}