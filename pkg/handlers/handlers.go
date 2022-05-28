package handlers

import "gorm.io/gorm"

//initialiser la connexion à la db en faisaint de l'injection de dependence
//pour ne pas initialiser la connexion dans chaque handler
//et creer pleins de connexion
type handler struct {
    DB *gorm.DB
}

func New(db *gorm.DB) handler {
    return handler{db}
}