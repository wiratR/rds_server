// utils/readfile.go
package utils

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile reads and prints the content of the file whose path is specified in the .env file
func ReadFileKey() ([]byte, error) {
	// Load environment variables from .env file
	// config, _ := config.LoadConfig(".")

	// // Get the file path from the environment variable
	// filePath := os.Getenv(config.FileKeyPath)
	// if filePath == "" {
	// 	return nil, fmt.Errorf("FILE_PATH environment variable not set")
	// }

	// Open the file
	filePath := "./config/key/Mch38806_PrivateKey.pem"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return content, nil
}
