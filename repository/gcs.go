package repository

import "context"

type Gcs interface {
	UploadFile(ctx context.Context, object string, fileBytes []byte) error
}
