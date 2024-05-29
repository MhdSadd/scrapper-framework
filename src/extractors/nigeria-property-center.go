package extractors

import (
	"encoding/json"
	"log"

	"github.com/MhdSadd/scrapper-framework/src/common"
	"github.com/MhdSadd/scrapper-framework/src/transformers"
	"github.com/MhdSadd/scrapper-framework/src/utils"
)

func NPCExtract(config common.Config) {
	// Use a map to store data for each URL
	dataStore := make(map[string][]map[string]string)
	transformedDataStore := make([]map[string]interface{}, 0)

	log.Println("⏳ mapping through extractors and getting detail URLs for scrapping...")
	for _, extractor := range config.Extractors {
		// activeSteps := extractor.ActiveSteps;
		detailPaths := Extract(extractor.Steps[0].Website, extractor.Steps[0].AttributePaths)

		for j, urls := range detailPaths {
			log.Println("⏳ Extracting data from detail path at index ", j)
			for key, value := range urls {
				if key != "source" {
					url, err := utils.ExtractPathFromURL(value)
					if err != nil {
						log.Fatal("❌ Failed to extract URL: ", err)
					}

					baseUrl, err := utils.ExtractBaseFromURL(extractor.Steps[0].Website)
					if err != nil {
						log.Fatal("❌ Failed to extract base URL: ", err)
					}

					url = baseUrl + url
					dataFromPage := Extract(url, extractor.Steps[1].AttributePaths)

					dataStore[url] = dataFromPage
				}
			}
		}

		transformedDataStore = transformers.Transform(dataStore)
	}

	encode, err := json.MarshalIndent(transformedDataStore, "", "    ")
	if err != nil {
		log.Fatal("❌ error indenting data\n", err)
	}

	utils.DataFileWriter(encode)
}
