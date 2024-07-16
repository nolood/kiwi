package main

import (
	"fmt"
	"kiwi/internal/config"
	"kiwi/internal/storage/postgres"
	"log"
)

func main() {
	cfg := config.MustLoad()

	storage := postgres.New(cfg.Storage)

	filePath := "/home/cities5000.txt"

	copyCommand := fmt.Sprintf("copy cities FROM '%s' WITH (FORMAT text, DELIMITER E'\\t', NULL '')", filePath)

	_, err := storage.Db.Exec(copyCommand)
	if err != nil {
		log.Fatalf("Failed to execute copy command: %v", err)
	}

	fmt.Println("Data imported successfully.")
}
