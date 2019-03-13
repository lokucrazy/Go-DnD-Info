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
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Println(w.Write(response))
			}
		})
		r.Get("/{table}/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
			response, err := get(chi.URLParam(r, "table"), name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Println(w.Write(response))
			}
		})
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
