package http_handlers

import (
	"net/http"

	"github.com/olebedev/emitter"
)

func ListenToEvents(eventEmitter *emitter.Emitter) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Add("Content-Type", "text/event-stream")

		print("Listening on events...")

		responseWriter.Write([]byte("event: stuff"))
		responseWriter.Write([]byte("\n"))
		responseWriter.Write([]byte("data: Welcome"))
		responseWriter.Write([]byte("\n\n"))

		for event := range eventEmitter.On("stuff") {
			println("New event!")

			responseWriter.Write([]byte("event: stuff"))
			responseWriter.Write([]byte("\n"))
			responseWriter.Write([]byte("data:"))
			responseWriter.Write([]byte(event.String(0)))
			responseWriter.Write([]byte("\n\n"))
		}
	}
}