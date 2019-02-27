package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/get", func(r chi.Router) {
		r.Get("/{table}", func(w http.ResponseWriter, r *http.Request) {
			j := json.NewEncoder(w)
			response, err := get(chi.URLParam(r, "table"), "")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				err = j.Encode(err)
			} else {
				w.WriteHeader(http.StatusOK)
				err = j.Encode(response)
			}
			fmt.Println(err)
		})
		r.Get("/{table}/{name}", func(w http.ResponseWriter, r *http.Request) {
			j := json.NewEncoder(w)
			name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
			response, err := get(chi.URLParam(r, "table"), name)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				err = j.Encode(err)
			} else {
				w.WriteHeader(http.StatusOK)
				err = j.Encode(response)
			}
		})
	})

	fmt.Println("listening on localhost:8081")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println(err)
	}
}
