package http_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/olebedev/emitter"

	"coso/gcs_client"
)

type GetRequestBody struct {
	Key string `json:"key"`
}

func GetValue(gcsClient *gcs_client.GCSClient, eventEmitter *emitter.Emitter) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var parsedRequestBody GetRequestBody

		jsonDecodingError := json.NewDecoder(request.Body).Decode(&parsedRequestBody)
		if jsonDecodingError != nil {
			println("Error decoding JSON:", jsonDecodingError.Error())
			http.Error(responseWriter, jsonDecodingError.Error(), http.StatusBadRequest)
			return
		}

		value, error := gcsClient.GetValue(parsedRequestBody.Key)
		if error != nil {
			println("Error getting value:", error.Error())
			http.Error(responseWriter, error.Error(), http.StatusBadRequest)
			return
		}

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(value)
	}
}
