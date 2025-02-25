package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/handlers/createChecklist"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/handlers/deleteChecklist"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/handlers/getChecklist"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/handlers/listChecklists"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/handlers/updateChecklist"
)

func New(checklistsRepository repositories.ChecklistsRepository) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", listChecklists.New(checklistsRepository).Handle)
	r.Post("/", createChecklist.New(checklistsRepository).Handle)
	r.Get("/{id}", getChecklist.New(checklistsRepository).Handle)
	r.Put("/{id}", updateChecklist.New(checklistsRepository).Handle)
	r.Delete("/{id}", deleteChecklist.New(checklistsRepository).Handle)

	return r
}
