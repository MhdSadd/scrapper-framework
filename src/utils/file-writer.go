package utils

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func DataFileWriter(data []byte) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := "scrapped_" + time.Now().Format("20060102150405") + ".json"
	filePath := filepath.Join(homeDir, "Downloads", fileName)

	log.Println("📝 writing data to file...")
	if err := os.WriteFile(filePath, data, 0664); err == nil {
		log.Println("✅ data written to file successfully")
	} else {
		log.Fatal(err)
	}
}
