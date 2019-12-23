package controllers

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type DnDInfoController interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
}

func CreateRouter(controller DnDInfoController) chi.Router {
	r := chi.NewRouter()
	r.Get("/", controller.GetAll)
	r.Get("/{name}", controller.GetByName)
	return r
}

func writeResponse(w http.ResponseWriter, response []byte, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println(w.Write(response))
	}
}
