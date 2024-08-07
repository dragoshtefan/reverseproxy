package filewriter

import (
	"log"
	"os"
	"path/filepath"
)

func WriteToFile(filePath string, content []byte) (int, error) {
	confDir := filepath.Dir(filePath)
	err := os.MkdirAll(confDir, os.ModePerm)

	if err != nil {
		log.Fatalf("Failed to create directory %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bytesWritten, err := file.Write(content)
	if err != nil {
		return 0, err
	}

	return bytesWritten, nil
}
