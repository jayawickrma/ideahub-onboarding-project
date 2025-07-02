package Controller

import (
	"Backend/DTO"
	"Backend/Service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type IdeaController struct{}

func NewIdeaController() *IdeaController {
	return &IdeaController{}
}

func (ic *IdeaController) CreateIdea(w http.ResponseWriter, r *http.Request) {
	var ideaDto DTO.IdeaDTO

	err := json.NewDecoder(r.Body).Decode(&ideaDto)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = Service.SaveIdea(ideaDto)
	if err != nil {
		http.Error(w, "Failed to save idea", http.StatusInternalServerError)
		return
	}
}

func (ic *IdeaController) DeleteIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := Service.RemoveIdea(id)
	if err != nil {
		http.Error(w, "Failed to delete idea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ic *IdeaController) UpdateIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var ideaDto DTO.IdeaDTO
	err := json.NewDecoder(r.Body).Decode(&ideaDto)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = Service.PutUpdate(ideaDto, id)
	if err != nil {
		http.Error(w, "Failed to update idea: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ic *IdeaController) GetAllIdeas(w http.ResponseWriter, r *http.Request) {
	err := Service.ViewAllIdeas()
	if err != nil {
		http.Error(w, "Error getting ideas", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("✅ Ideas fetched successfully")) // or return data
}
