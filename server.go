package main

import (
	"database/sql"
	"log"
	"lokucrazy/Go-DnD-Info/controllers"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("port not provided")
	}
	port := os.Args[1]
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatal("port must be a valid number")
	}

	db, err := sql.Open("sqlite3", "5eInfo.db")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer closeConnection(db)

	r := chi.NewRouter()
	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cor.Handler)

	r.Mount("/spells", controllers.CreateRouter(controllers.CreateSpellsController(db)))
	r.Mount("/weapons", controllers.CreateRouter(controllers.CreateWeaponsController(db)))
	r.Mount("/armors", controllers.CreateRouter(controllers.CreateArmorsController(db)))
	r.Mount("/levels", controllers.CreateRouter(controllers.CreateLevelsController(db)))
	r.Mount("/proficiencies", controllers.CreateRouter(controllers.CreateProficienciesController(db)))
	r.Mount("/races", controllers.CreateRouter(controllers.CreateRacesController(db)))
	r.Mount("/feats", controllers.CreateRouter(controllers.CreateFeatsController(db)))

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func closeConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
