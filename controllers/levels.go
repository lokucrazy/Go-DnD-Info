package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type levelsController struct {
	service services.Service
}

func (l *levelsController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := l.service.Get("")
	writeResponse(w, response, err)
}

func (l *levelsController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := l.service.Get(name)
	writeResponse(w, response, err)
}

func CreateLevelsController(db *sql.DB) *levelsController {
	return &levelsController{service: services.CreateLevelsService(db)}
}
