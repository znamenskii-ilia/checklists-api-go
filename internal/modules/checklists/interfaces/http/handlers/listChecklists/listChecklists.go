package listChecklists

import (
	"net/http"

	"github.com/znamenskii-ilia/checklists-api-go/internal/buildingBlocks/httpUtils"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/infrastructure/repositories"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/dtos"
	"github.com/znamenskii-ilia/checklists-api-go/internal/modules/checklists/interfaces/http/mappers"
)

type ListChecklistsHandler struct {
	checklistsRepository repositories.ChecklistsRepository
}

func New(checklistsRepository repositories.ChecklistsRepository) *ListChecklistsHandler {
	return &ListChecklistsHandler{checklistsRepository: checklistsRepository}
}

func (h *ListChecklistsHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cs, err := h.checklistsRepository.List()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	csDtos := []*dtos.ChecklistDTO{}

	for _, c := range cs {
		csDtos = append(csDtos, mappers.ToDTO(c))
	}

	httpUtils.WriteJSON(w, http.StatusOK, csDtos)
}
