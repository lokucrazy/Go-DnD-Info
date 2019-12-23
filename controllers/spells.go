package controllers

import (
	"database/sql"
	"github.com/go-chi/chi"
	"lokucrazy/Go-DnD-Info/services"
	"net/http"
	"strings"
)

type spellsController struct {
	service services.Service
}

func (s *spellsController) GetAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response, err := s.service.Get("")
	writeResponse(w, response, err)
}

func (s *spellsController) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
	response, err := s.service.Get(name)
	writeResponse(w, response, err)
}

func CreateSpellsController(db *sql.DB) *spellsController {
	return &spellsController{service: services.CreateSpellsService(db)}
}
