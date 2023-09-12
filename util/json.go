package util

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func ValidateImportFile(path string) bool {
	// Check if file exists
	stat, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	// Check if file is a directory
	if stat.IsDir() {
		return false
	}

	//Check if file is a valid json file
	if filepath.Ext(path) != ".json" {
		return false
	}

	// Check the contents and validate the json structure
	if err := validateJSON(path); err != nil {
		return false
	}

	return true
}

func validateJSON(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if !json.Valid(data) {
		return errors.New("invalid JSON")
	}

	return nil
}
