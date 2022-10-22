package main

import (
	"log"
	"net/http"
	"os"

	"github.com/olebedev/emitter"

	"coso/gcs_client"
	"coso/http_handlers"
)

func main() {
	eventEmitter := &emitter.Emitter{}

	bucketName := os.Getenv("BUCKET_NAME")
	portNumber := os.Getenv("PORT_NUMBER")
	if portNumber == "" {
		log.Printf("`PORT_NUMBER` environment variable not set, falling back to port 1337")
		portNumber = "1337"
	}

	if bucketName == "" {
		log.Printf("`BUCKET_NAME` environment variable is not set")
		os.Exit(1)
	}

	log.Printf("Bucket name: %s", bucketName)

	gcsClient, error := gcs_client.NewGCSClient(bucketName, eventEmitter)
	if error != nil {
		panic(error)
	}

	http.HandleFunc("/listen", http_handlers.ListenToEvents(eventEmitter))
	http.HandleFunc("/get", http_handlers.GetValue(gcsClient, eventEmitter))
	http.HandleFunc("/set", http_handlers.SetValue(gcsClient, eventEmitter))
	http.HandleFunc("/list", http_handlers.ListValues(gcsClient, eventEmitter))

	done := make(chan bool)
	go http.ListenAndServe(":"+portNumber, nil)
	log.Printf("Server listening on port %v...", portNumber)
	<-done
}
