package gcs_client

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"github.com/olebedev/emitter"
)

type GCSClient struct {
	// Client is the underlying GCS client.
	Client *storage.Client
	// Bucket is the GCS bucket to use.
	Bucket *storage.BucketHandle
	// BucketName is the name of the GCS bucket to use.
	BucketName string
	// EventEmitter is the event emitter to use.
	EventEmitter *emitter.Emitter
}

// NewGCSClient creates a new GCS client.
func NewGCSClient(bucketName string, eventEmitter *emitter.Emitter) (*GCSClient, error) {
	client, error := storage.NewClient(context.Background())
	if error != nil {
		return nil, error
	}

	return &GCSClient{
		Client:       client,
		Bucket:       client.Bucket(bucketName),
		BucketName:   bucketName,
		EventEmitter: eventEmitter,
	}, nil
}

func (gcsClient *GCSClient) SetValue(key string, value string) error {
	object := gcsClient.Bucket.Object(key)
	operationContext := context.Background()

	// Instantiate a new Writer and use it to write stuff.
	objectWriter := object.NewWriter(operationContext)

	// Write the value.
	if _, writeError := fmt.Fprintf(objectWriter, value); writeError != nil {
		return writeError
	}

	// Close, just like writing a file.
	if objectWriterCloseError := objectWriter.Close(); objectWriterCloseError != nil {
		return objectWriterCloseError
	}

	return nil
}

func (gcsClient *GCSClient) GetValue(key string) ([]byte, error) {
	object := gcsClient.Bucket.Object(key)
	operationContext := context.Background()

	// Read it back.
	objectReader, objectReaderError := object.NewReader(operationContext)
	if objectReaderError != nil {
		return nil, objectReaderError
	}

	content, ioError := ioutil.ReadAll(objectReader)
	if ioError != nil {
		return nil, ioError
	}

	// Close the reader.
	defer objectReader.Close()

	return content, nil

}
