// Package collections serves as a way to parse JSON into the UI.
// We'll receive our JSON with saved requests and use it as a file tree with additional info.
package collections

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
* Sample JSON:
* {
    "directory_name": {
      "request_name": {
        "method": "POST",
        "url": "localhost:3000/api/hello_world",
        "headers": {
          "Content-Type": "application/json",
          "Authorization": "Bearer {{token}}"
        },
        "body": "{\"key\": \"value\"}",
        "queryParams": {
          "limit": "10"
        }
      }
    }
  }
*/

type Request struct {
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Headers     map[string]string `json:"headers"`
	Body        string            `json:"body"`
	QueryParams map[string]string `json:"query_params"`
}

type Directory map[string]Request
type Collection map[string]Directory

func buildFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to locate home directory: %w", err)
	}
	configFile := homeDir + "/.config/go-postman/config.json"

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return "", fmt.Errorf("config file does not exist: %w", err)
	}

	return configFile, nil
}

func LoadFile() (*Collection, error) {
	configFile, err := buildFilePath()
	if err != nil {
		return nil, fmt.Errorf("cannot build file path: %w", err)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error opening config file: %w", err)
	}

	var collection Collection
	err = json.Unmarshal(data, &collection)
	if err != nil {
		return nil, fmt.Errorf("cannot read json config contents: %w", err)
	}

	return &collection, nil
}

func PrettyPrint(coll *Collection) (string, error) {
	pp, err := json.MarshalIndent(coll, "", "  ")
	if err != nil {
		return "", fmt.Errorf("unable to pretty print json: %w", err)
	}

	return string(pp), nil
}

func AddDirectory(coll *Collection, directory string) (*Collection, error) {

	if _, exists := (*coll)[directory]; exists {
		return nil, fmt.Errorf("directory already exists: %s", directory)
	}

	(*coll)[directory] = Directory{}

	filePath, err := buildFilePath()
	if err != nil {
		return nil, fmt.Errorf("cannot build file path for new directory: %w", err)
	}

	updatedJSON, err := json.MarshalIndent(coll, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("cannot update json w/ new directory: %w", err)
	}

	err = os.WriteFile(filePath, updatedJSON, 0644)
	if err != nil {
		return nil, fmt.Errorf("cannot write new directory to file: %w", err)
	}

	return coll, nil

}
