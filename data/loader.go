package data

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// LoadPrices reads float values from a given file path and returns a slice of floats.
func LoadPrices(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening the file: %w", err)
	}
	defer file.Close()

	var prices []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing float from line '%s': %w", line, err)
		}
		prices = append(prices, price)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %w", err)
	}

	return prices, nil
}
