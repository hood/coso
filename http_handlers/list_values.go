package http_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/olebedev/emitter"
	"google.golang.org/api/iterator"

	"coso/gcs_client"
)

type ListRequestBody struct {
	Prefix string `json:"prefix"`
	Flat   bool   `json:"flat"`
}

func ListValues(gcsClient *gcs_client.GCSClient, eventEmitter *emitter.Emitter) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var parsedRequestBody ListRequestBody

		jsonDecodingError := json.NewDecoder(request.Body).Decode(&parsedRequestBody)
		if jsonDecodingError != nil {
			println("Error decoding JSON:", jsonDecodingError.Error())

			if jsonDecodingError.Error() != "EOF" {
				http.Error(responseWriter, jsonDecodingError.Error(), http.StatusBadRequest)
				return
			} else {
				parsedRequestBody.Prefix = ""
			}
		}

		resultsDelimiter := "/"
		if parsedRequestBody.Flat != true {
			resultsDelimiter = ""
		}

		// Prefixes and delimiters can be used to emulate directory listings.
		// Prefixes can be used to filter objects starting with prefix.
		// The delimiter argument can be used to restrict the results to only the
		// objects in the given "directory". Without the delimiter, the entire tree
		// under the prefix is returned.
		//
		// For example, given these blobs:
		//   /a/1.txt
		//   /a/b/2.txt
		//
		// If you just specify prefix="a/", you'll get back:
		//   /a/1.txt
		//   /a/b/2.txt
		//
		// However, if you specify prefix="a/" and delim="/", you'll get back:
		//   /a/1.txt
		query := &storage.Query{
			Prefix:     parsedRequestBody.Prefix,
			Versions:   false,
			Projection: storage.ProjectionNoACL,
			Delimiter:  resultsDelimiter,
		}
		query.SetAttrSelection([]string{"Prefix", "Name"})

		itemsIterator := gcsClient.Bucket.Objects(context.Background(), query)

		var paths []string

		for {
			currentItem, itemRetrievalError := itemsIterator.Next()
			if itemRetrievalError != nil {
				if itemRetrievalError == iterator.Done {
					break
				}

				fmt.Println(itemRetrievalError)
				break
			}

			paths = append(paths, currentItem.Prefix+currentItem.Name)
		}

		var items []any

		itemsChannel := make(chan string)

		for _, path := range paths {
			currentPath := path

			go func(channel chan string, objectPath string, gcsClient *gcs_client.GCSClient, responseWriter http.ResponseWriter) {
				item, error := gcsClient.GetValue(objectPath)
				if error != nil {
					println("Error getting value:", error.Error())
					http.Error(responseWriter, error.Error(), http.StatusBadRequest)
					return
				}

				channel <- string(item)
			}(itemsChannel, currentPath, gcsClient, responseWriter)
		}

		for retrievedItem := range itemsChannel {
			items = append(items, retrievedItem)

			if len(items) >= len(paths) {
				break
			}
		}

		serializedResults, error := json.Marshal(items)
		if error != nil {
			println("Error serializing results:", error.Error())
			http.Error(responseWriter, error.Error(), http.StatusBadRequest)
			return
		}

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(serializedResults))
	}
}
