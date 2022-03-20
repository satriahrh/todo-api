package dto

type CreateItemRequest struct {
	Name        string
	Description string
}

type CreateItemWithFileRequest struct {
	CreateItemRequest
	Document []byte
}

type UpdateItemRequest struct {
	ID          int
	Name        string
	Description string
}
