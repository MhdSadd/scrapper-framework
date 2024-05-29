package main

import (
	"log"
	"os"

	"github.com/MhdSadd/scrapper-framework/src/configurations"
	"github.com/MhdSadd/scrapper-framework/src/extractors"
)

func main() {
	log.Println("Scraper Framework Initialized 🚀🚀🚀")

	// Check if a file path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Println("ℹ️ please specify a file path. Usage: go run main.go /Desktop/config.yml")
		os.Exit(1)
	}

	// Get the file path from command-line arguments
	configFilePath := os.Args[1]

	config := configurations.ValidateConfigFile(configFilePath)

	extractors.NPCExtract(config)
}
