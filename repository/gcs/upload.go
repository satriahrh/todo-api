package gcs

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type gcs struct {
}

func New() *gcs {
	return &gcs{}
}

// uploadFile uploads an object.
func (*gcs) UploadFile(ctx context.Context, object string, fileBytes []byte) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.

	wc := client.Bucket(os.Getenv("GCS_BUCKET")).Object(object).NewWriter(ctx)

	_, err = wc.Write(fileBytes)

	if err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}
	return nil
}
