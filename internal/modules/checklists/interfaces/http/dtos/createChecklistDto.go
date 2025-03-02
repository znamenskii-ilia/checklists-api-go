package dtos

type CreateChecklistDto struct {
	Title string `json:"title" validate:"required"`
}
