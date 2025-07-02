package Controller

import (
	"Backend/DTO"
	"Backend/Service"
	"encoding/json"
	"fmt"
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
	fmt.Fprintf(w, "❌ DeleteIdea called for ID: %s\n", id)
}

func (ic *IdeaController) UpdateIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "✏️ UpdateIdea called for ID: %s\n", id)
}

func (ic *IdeaController) GetAllIdeas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "📥 GetAllIdeas called")
}
