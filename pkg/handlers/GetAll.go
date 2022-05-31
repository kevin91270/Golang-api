package handlers

import (
    "encoding/json"
    "net/http"
    "fmt"

    "github.com/kevin91270/Golang-api/pkg/models"
)

//ajout de h handler avant le nom de la fonction
//pour la definir en receiver function
//getallfilms est une fonction qui peut recevoir un handler
//de connexion afin d'eviter de creer une nouvelle connexion a la db
//obtenir les infos de tout les films
func (h handler) GetAllFilms(w http.ResponseWriter, r *http.Request) {
    var films []models.Film

    if result := h.DB.Find(&films); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(films)
}

//obtenir les infos de tout les users
func (h handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User

    if result := h.DB.Preload("Films").Find(&users); result.Error != nil {
        fmt.Println(result.Error)
    }

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}