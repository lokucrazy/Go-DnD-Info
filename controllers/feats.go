package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type featsController struct {
	service services.Service
}

func (f *featsController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := f.service.Get("")
	writeResponse(w, response, err)
}

func (f *featsController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := f.service.Get(name)
	writeResponse(w, response, err)
}

func CreateFeatsController(db *sql.DB) *featsController {
	return &featsController{service: services.CreateFeatsService(db)}
}
