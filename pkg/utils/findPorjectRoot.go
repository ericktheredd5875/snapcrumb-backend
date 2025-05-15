package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindProjectRoot(startDir string) (string, error) {

	var err error
	curDir := startDir
	if startDir == "" {
		curDir, err = os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %v", err)
		}
	}

	for {
		if _, err := os.Stat(filepath.Join(curDir, "go.mod")); err == nil {
			curDir = filepath.ToSlash(curDir)
			return curDir, nil
		}

		parentDir := filepath.Dir(curDir)
		if parentDir == curDir {
			return "", fmt.Errorf("project root not found")
		}

		curDir = parentDir
	}

}
