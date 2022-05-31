package handlers

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "fmt"

    "github.com/gorilla/mux"
    "github.com/kevin91270/Golang-api/pkg/models"
)

//ajout de h handler avant le nom de la fonction
//pour la definir en receiver function
//updatefilm est une fonction qui peut recevoir un handler
//de connexion afin d'eviter de creer une nouvelle connexion a la db
//mettre à jour un film
func (h handler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
    // lire un id dynamique en parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // lire la requete dans le body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    //parse JSON
    var updatedFilm models.Film
    json.Unmarshal(body, &updatedFilm)

    var film models.Film

    if result := h.DB.First(&film, id); result.Error != nil {
        fmt.Println(result.Error)
    }
    //mise a jour des champs
    film.Title = updatedFilm.Title
    film.Director = updatedFilm.Director
    film.Synopsis = updatedFilm.Synopsis
    film.Score = updatedFilm.Score
    film.UserId = updatedFilm.UserId

    //sauvegarde dans la db
    h.DB.Save(&film)

    //send response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}

//mettre à jour un user
func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
    // lire un id dynamique en parametre
    vars := mux.Vars(r)
    id, _ := strconv.Atoi(vars["id"])

    // lire la requete dans le body
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        log.Fatalln(err)
    }

    //parse JSON
    var updatedUser models.User
    json.Unmarshal(body, &updatedUser)

    var user models.User

    if result := h.DB.First(&user, id); result.Error != nil {
        fmt.Println(result.Error)
    }
    //mise a jour des champs
    user.Pseudo = updatedUser.Pseudo

    //sauvegarde dans la db
    h.DB.Save(&user)

    //send response
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Updated")
}