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
    router.HandleFunc("/films/{id}", h.GetFilm).Methods(http.MethodGet)
    router.HandleFunc("/films", h.AddFilm).Methods(http.MethodPost)
    router.HandleFunc("/films/{id}", h.UpdateFilm).Methods(http.MethodPut)
    router.HandleFunc("/films/{id}", h.DeleteFilm).Methods(http.MethodDelete)

	router.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
    router.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
    router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
    router.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPut)
    router.HandleFunc("/users/{id}", h.DeleteUser).Methods(http.MethodDelete)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}