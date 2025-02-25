package createChecklist

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/domainErrors"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/httpUtils"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/domain"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/dtos"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/mappers"
)

type CreateChecklistHandler struct {
	checklistsRepository repositories.ChecklistsRepository
}

func New(checklistsRepository repositories.ChecklistsRepository) *CreateChecklistHandler {
	return &CreateChecklistHandler{checklistsRepository: checklistsRepository}
}

func (h *CreateChecklistHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var dto dtos.CreateChecklistDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := &domain.Checklist{
		ID:    uuid.New().String(),
		Title: dto.Title,
		Tasks: []domain.Task{},
	}

	c, err = h.checklistsRepository.CreateOne(c)
	if err != nil {
		if errors.Is(err, domainErrors.ErrEntityConflict) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpUtils.WriteJSON(w, http.StatusOK, mappers.ToDTO(c))
}
