package models

//model film
//utilisation de gorm pour les drivers postgres 
//ajout du tag cle primaire sur l'ID
type Film struct {
    Id     int    `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
    Desc   string `json:"desc"`
}