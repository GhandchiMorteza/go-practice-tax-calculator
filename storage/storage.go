package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// SaveToJSON takes a generic data interface{}, marshals it to JSON, and saves it to the specified file path.
func SaveToJSON(data interface{}, name string) error {
		file, err := os.Create(name)
		if err != nil {
			return fmt.Errorf("failed to create JSON file: %w", err)
		}
		defer file.Close()

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		
		if err := encoder.Encode(data); err != nil {
			return fmt.Errorf("JSON encoding failed: %w", err)
		}
		
    return nil
}
