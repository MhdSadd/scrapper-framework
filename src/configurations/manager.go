package configurations

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/MhdSadd/scrapper-framework/src/common"
)

func ValidateConfigFile(filePath string) common.Config {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("❌ error reading YAML file: %v", err)
	}

	// Unmarshal YAML data into Config struct
	var config common.Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("❌ error unmarshaling YAML: %v", err)
	}

	if len(config.Extractors) == 0 {
		log.Fatalf("❌ Extractors cannot be empty: atleast one extractor is required to preceed")
		os.Exit(1)
	}

	for _, extractor := range config.Extractors {
		for _, step := range extractor.Steps {
			if step.Number < 1 {
				log.Fatalf("❌ Step number cannot be empty: please provide an integer value starting from 1")
				os.Exit(1)
			}

			if len(step.AttributePaths) == 0 {
				log.Fatalf("❌ attributePaths cannot be empty: atleast one attribute path is required to preceed")
				os.Exit(1)
			}
		}
	}

	return config
}
