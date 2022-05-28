package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/kevin91270/Golang-api/pkg/db"
    "github.com/kevin91270/Golang-api/pkg/handlers"
)

func main() {
	//initialisation de la connexion a la db
    DB := db.Init()
    h := handlers.New(DB)
	//utilisation de mux pour le routage
    router := mux.NewRouter()

	//differente route pour l'api
    router.HandleFunc("/films", h.GetAllFilms).Methods(http.MethodGet)
    router.HandleFunc("/film/{id}", h.GetFilm).Methods(http.MethodGet)
    router.HandleFunc("/film", h.AddFilm).Methods(http.MethodPost)
    router.HandleFunc("/film/{id}", h.UpdateFilm).Methods(http.MethodPut)
    router.HandleFunc("/film/{id}", h.DeleteFilm).Methods(http.MethodDelete)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}