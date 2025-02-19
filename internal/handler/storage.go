package handler

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

const storageFile = "uploads\\data\\storage.csv"

func SaveFileInfo(productName, fileName string) error {
	file, err := os.OpenFile(storageFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open storage file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{productName, fileName})
	if err != nil {
		return fmt.Errorf("failed to write to storage file: %v", err)
	}
	return nil
}

func UploadToTripleS(fileData []byte, fileName string) (string, error) {
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	filePath := filepath.Join(uploadDir, fileName)
	err := os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	err = SaveFileInfo(fileName, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to save file info: %v", err)
	}

	return filePath, nil
}

func DeleteFromTripleS(fileName string) error {
	filePath := filepath.Join("uploads", fileName)
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %v", err)
	}

	return nil
}
