package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type armorsController struct {
	service services.Service
}

func (a *armorsController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := a.service.Get("")
	writeResponse(w, response, err)
}

func (a *armorsController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := a.service.Get(name)
	writeResponse(w, response, err)
}

func CreateArmorsController(db *sql.DB) *armorsController {
	return &armorsController{service: services.CreateArmorsService(db)}
}
