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

		value, getValueError := gcsClient.GetValue(parsedRequestBody.Key)
		if getValueError != nil {
			println("Error getting value:", getValueError.Error())
			http.Error(responseWriter, getValueError.Error(), http.StatusBadRequest)
			return
		}

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write(value)
	}
}
