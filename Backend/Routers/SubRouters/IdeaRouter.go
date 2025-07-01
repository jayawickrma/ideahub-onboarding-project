package routes

import (
	"Backend/Controller"
	"github.com/gorilla/mux"
)

type IdeaRoutes struct {
	Router         *mux.Router
	IdeaController *Controller.IdeaController
}

func NewIdeaRoutes() *IdeaRoutes {
	r := mux.NewRouter()
	ic := Controller.NewIdeaController()

	ideaRoutes := &IdeaRoutes{
		Router:         r,
		IdeaController: ic,
	}

	ideaRoutes.initRoutes()
	return ideaRoutes
}

func (ir *IdeaRoutes) initRoutes() {
	ir.Router.HandleFunc("/view", ir.IdeaController.GetAllIdeas).Methods("GET")
	ir.Router.HandleFunc("/createIdea", ir.IdeaController.CreateIdea).Methods("POST")
	ir.Router.HandleFunc("/updateIdea", ir.IdeaController.UpdateIdea).Methods("PUT")
	ir.Router.HandleFunc("/deleteIdea", ir.IdeaController.DeleteIdea).Methods("DELETE")
}
