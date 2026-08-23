package filehandling

import (
	"fmt"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("failed to read the input file %q: %w", path, err)
	}

	return data, nil
}
