package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/kevin91270/Golang-api/pkg/models"
)

//ajout de h handler avant le nom de la fonction
//pour la definir en receiver function
//getfilm est une fonction qui peut recevoir un handler
//de connexion afin d'eviter de creer une nouvelle connexion a la db
func (h handler) GetFilm(w http.ResponseWriter, r *http.Request) {
    // lire un id dynamique en parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // trouver un film avec l'Id
    var film models.Film

    if result := h.DB.First(&film, id); result.Error != nil {
        fmt.Println(result.Error)
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(film)
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
    // lire un id dynamique en parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // trouver un un avec l'Id
    var user models.User

    if result := h.DB.Preload("Films").First(&user, id); result.Error != nil {
        fmt.Println(result.Error)
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}