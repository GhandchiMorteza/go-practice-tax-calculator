package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) Process() {
	result := map[string]string{}

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	fmt.Println(result)
}
func loadingData() ([]float64, error) {
	file, err := os.Open("prices/prices.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %w", err)
	}
	defer file.Close()

	var floats []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		f, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing float from line '%s': %w", line, err)
		}
		floats = append(floats, f)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return floats, nil
}

func NewTaxIncludedPriceJob(taxRate float64) (*TaxIncludedPriceJob, error) {
	prices, err := loadingData()
	if err != nil {
		return nil, err
	}
	return &TaxIncludedPriceJob{
		InputPrices: prices,
		TaxRate:     taxRate,
	}, nil
}