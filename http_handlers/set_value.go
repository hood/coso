package http_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/olebedev/emitter"

	"coso/gcs_client"
)

type SetRequestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func SetValue(gcsClient *gcs_client.GCSClient, eventEmitter *emitter.Emitter) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var parsedRequestBody SetRequestBody

		jsonDecodingError := json.NewDecoder(request.Body).Decode(&parsedRequestBody)
		if jsonDecodingError != nil {
			http.Error(responseWriter, jsonDecodingError.Error(), http.StatusBadRequest)
			return
		}

		setValueError := gcsClient.SetValue(parsedRequestBody.Key, parsedRequestBody.Value)
		if setValueError != nil {
			http.Error(responseWriter, setValueError.Error(), http.StatusBadRequest)
			return
		}

		eventEmitter.Emit(parsedRequestBody.Key, parsedRequestBody.Value)

		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Write([]byte(parsedRequestBody.Value))
	}
}
