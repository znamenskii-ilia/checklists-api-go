package updateChecklist

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/httpUtils"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/dtos"
)

type UpdateChecklistHandler struct {
	checklistsRepository repositories.ChecklistsRepository
}

func New(checklistsRepository repositories.ChecklistsRepository) *UpdateChecklistHandler {
	return &UpdateChecklistHandler{checklistsRepository: checklistsRepository}
}

func (h *UpdateChecklistHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "id")

	var dto dtos.ChecklistDTO
	err := httpUtils.DecodeBody(r, &dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := &dtos.ChecklistDTO{
		ID:    cID,
		Title: dto.Title,
		Tasks: dto.Tasks,
	}

	httpUtils.WriteJSON(w, http.StatusOK, c)
}
