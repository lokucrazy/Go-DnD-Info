package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("port not provided")
	}
	port := os.Args[1]
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatal("port must be a valid number")
	}

	r := chi.NewRouter()
	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cor.Handler)

	r.Route("/get", func(r chi.Router) {

		r.Get("/{table}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			response, err := get(chi.URLParam(r, "table"), "")
			writeResponse(w, response, err)
		})

		r.Get("/{table}/{name}", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			name := strings.Replace(chi.URLParam(r, "name"), "_", " ", -1)
			response, err := get(chi.URLParam(r, "table"), name)
			writeResponse(w, response, err)
		})
	})

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func writeResponse(w http.ResponseWriter, b []byte, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Println(w.Write(b))
	}
}
