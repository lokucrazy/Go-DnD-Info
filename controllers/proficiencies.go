package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type proficienciesController struct {
	service services.Service
}

func (p *proficienciesController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := p.service.Get("")
	writeResponse(w, response, err)
}

func (p *proficienciesController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := p.service.Get(name)
	writeResponse(w, response, err)
}

func CreateProficienciesController(db *sql.DB) *proficienciesController {
	return &proficienciesController{service: services.CreateProficienciesService(db)}
}
