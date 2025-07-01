package Routers

import (
	routes "Backend/Routers/SubRouters"
	"github.com/gorilla/mux"
	"net/http"
)

type MainRouter struct {
	Router     *mux.Router
	IdeaRouter *routes.IdeaRoutes
}

func NewRouter() *MainRouter {
	r := mux.NewRouter()
	ir := routes.NewIdeaRoutes()

	mainRoutes := &MainRouter{
		Router:     r,
		IdeaRouter: ir,
	}

	mainRoutes.initRoutes()
	return mainRoutes
}

func (m *MainRouter) initRoutes() {
	m.Router.PathPrefix("/idea").Handler(m.IdeaRouter.Router)
	m.Router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is up and running"))
	})
}
