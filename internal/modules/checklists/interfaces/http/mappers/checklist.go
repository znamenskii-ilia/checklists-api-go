package mappers

import (
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/domain"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/dtos"
)

func ToDTO(checklist *domain.Checklist) *dtos.ChecklistDTO {
	return &dtos.ChecklistDTO{
		ID:    checklist.ID,
		Title: checklist.Title,
		Tasks: ToDTOTasks(checklist.Tasks),
	}
}

func ToDTOTasks(tasks []domain.Task) []dtos.TaskDTO {
	dtos := make([]dtos.TaskDTO, len(tasks))

	for i, task := range tasks {
		dtos[i] = ToDTOTask(task)
	}

	return dtos
}

func ToDTOTask(task domain.Task) dtos.TaskDTO {
	return dtos.TaskDTO{
		Title: task.Title,
	}
}
