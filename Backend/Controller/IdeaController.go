package Controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type IdeaController struct{}

func NewIdeaController() *IdeaController {
	return &IdeaController{}
}

func (ic *IdeaController) CreateIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🆕 CreateIdea called")
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
