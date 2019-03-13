package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/get", func(r chi.Router) {

		r.Get("/{table}", func(w http.ResponseWriter, r *http.Request) {
			response, err := get(chi.URLParam(r, "table"), "")
			writeResponse(w, response, err)
		})

		r.Get("/{table}/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
			response, err := get(chi.URLParam(r, "table"), name)
			writeResponse(w, response, err)
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}

func writeResponse(w http.ResponseWriter, b []byte, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Println(w.Write(b))
	}
}
