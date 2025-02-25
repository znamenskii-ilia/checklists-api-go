package getChecklist

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/domainErrors"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/httpUtils"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/mappers"
)

type GetChecklistHandler struct {
	checklistsRepository repositories.ChecklistsRepository
}

func New(checklistsRepository repositories.ChecklistsRepository) *GetChecklistHandler {
	return &GetChecklistHandler{checklistsRepository: checklistsRepository}
}

func (h *GetChecklistHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "id")

	c, err := h.checklistsRepository.FindOne(cID)

	if err != nil {
		if errors.Is(err, domainErrors.ErrEntityNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dto := mappers.ToDTO(c)

	httpUtils.WriteJSON(w, http.StatusOK, dto)
}
