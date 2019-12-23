package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type weaponsController struct {
	service services.Service
}

func (we *weaponsController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := we.service.Get("")
	writeResponse(w, response, err)
}

func (we *weaponsController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := we.service.Get(name)
	writeResponse(w, response, err)
}

func CreateWeaponsController(db *sql.DB) *weaponsController {
	return &weaponsController{service: services.CreateWeaponsService(db)}
}
