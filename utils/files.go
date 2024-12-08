package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFromFile(filename string, extension string, fileDir string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %v", err)
	}

	filePath := filepath.Join(currentDir, "aoc", fileDir, filename+extension)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %v", filename, err)
	}
	return string(content), nil
}
