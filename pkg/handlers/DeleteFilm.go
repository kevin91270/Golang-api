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
//deleteFilm est une fonction qui peut recevoir un handler
//de connexion afin d'eviter de creer une nouvelle connexion a la db
func (h handler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
    // lire un id dynamique en parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // trouver le film avec l'Id

    var film models.Film

    if result := h.DB.First(&film, id); result.Error != nil {
        fmt.Println(result.Error)
    }

    // suppression du film
    h.DB.Delete(&film)

    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Deleted")
}