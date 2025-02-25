package dtos

type ChecklistDTO struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Tasks []TaskDTO `json:"tasks"`
}

type TaskDTO struct {
	Title string `json:"title"`
}
