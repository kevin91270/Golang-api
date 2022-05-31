package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	"github.com/kevin91270/Golang-api/pkg/models"
)

//ajout de h handler avant le nom de la fonction
//pour la definir en receiver function
//addfilm est une fonction qui peut recevoir un handler
//de connexion afin d'eviter de creer une nouvelle connexion a la db
//fonction pour ajouter un film
func (h handler) AddFilm(w http.ResponseWriter, r *http.Request) {
	// lire la requete
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	var film models.Film
	json.Unmarshal(body, &film)

	// ajouter à la table Films
	if result := h.DB.Create(&film); result.Error != nil {
		fmt.Println(result.Error)
	}

	// retourne code de reponse 201 creation
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

//fonction pour ajouter un user
func (h handler) AddUser(w http.ResponseWriter, r *http.Request) {
	// lire la requete
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	var user models.User
	json.Unmarshal(body, &user)

	// ajouter à la table Users
	if result := h.DB.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}

	// retourne code de reponse 201 creation
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}