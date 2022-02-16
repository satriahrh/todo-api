package dto

type CreateItemRequest struct {
	Name        string
	Description string
}

type UpdateItemRequest struct {
	ID          int
	Name        string
	Description string
}
