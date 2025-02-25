package repositories

import "github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/domain"

type ChecklistsRepository interface {
	List() ([]*domain.Checklist, error)
	CreateOne(checklist *domain.Checklist) (*domain.Checklist, error)
	FindOne(id string) (*domain.Checklist, error)
	SaveOne(checklist *domain.Checklist) (*domain.Checklist, error)
	DeleteOne(id string) error
}
