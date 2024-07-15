package main

import (
	"fmt"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"
	"log"
	"os"
	"path/filepath"
)

func main() {
	cfg := config.MustLoad()

	storage := postgres.New(cfg.Storage)

	fileName := "cities5000.txt"
	absFilePath, err := filepath.Abs(filepath.Join("./static", fileName))
	if err != nil {
		log.Fatalf("Failed to get absolute path for file: %v", err)
	}

	if _, err = os.Stat(absFilePath); os.IsNotExist(err) {
		panic("data file does not exist: " + absFilePath)
	}

	copyCommand := fmt.Sprintf("copy cities FROM '%s' WITH (FORMAT text, DELIMITER E'\\t', NULL '')", absFilePath)

	_, err = storage.Db.Exec(copyCommand)
	if err != nil {
		log.Fatalf("Failed to execute copy command: %v", err)
	}

	fmt.Println("Data imported successfully.")
}
