package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type racesController struct {
	service services.Service
}

func (ra *racesController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := ra.service.Get("")
	writeResponse(w, response, err)
}

func (ra *racesController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := ra.service.Get(name)
	writeResponse(w, response, err)
}

func CreateRacesController(db *sql.DB) *racesController {
	return &racesController{service: services.CreateRacesService(db)}
}
