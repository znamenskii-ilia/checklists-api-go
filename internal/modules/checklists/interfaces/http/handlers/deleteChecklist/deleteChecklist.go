package deleteChecklist

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/httpUtils"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
)

type DeleteChecklistHandler struct {
	checklistsRepository repositories.ChecklistsRepository
}

func New(checklistsRepository repositories.ChecklistsRepository) *DeleteChecklistHandler {
	return &DeleteChecklistHandler{checklistsRepository: checklistsRepository}
}

func (h *DeleteChecklistHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cID := chi.URLParam(r, "id")

	err := h.checklistsRepository.DeleteOne(cID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpUtils.WriteJSON(w, http.StatusOK, nil)
}
