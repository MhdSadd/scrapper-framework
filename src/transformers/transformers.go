package transformers

import (
	"log"
)

func Transform( data map[string][]map[string]string) []map[string] interface{} {
	// value is the array of objects
	log.Println("⏳ Transforming data...")
	dataStore := make([]map[string] interface{}, 0)
	for key, value := range data {
        tempStore := make(map[string]interface{})

        // check if each key is duplicated in other sub-objects
        for i := 0; i < len(value); i++ {
            for k, v := range value[i] {
                val, exists := tempStore[k]
                // if it exists and is not a duplicate
                if exists && v != tempStore[k] {
                    if arr, isArray := val.([]string); isArray {
                        tempStore[k] = append(arr, v)
                    } else {
                        tempStore[k] = []string{val.(string), v}
                    }
                } else {
                    tempStore[k] = v
                }
            }
        }

       tempStore["url"] = key; 
	   dataStore = append(dataStore, tempStore);
    }

	return dataStore
}