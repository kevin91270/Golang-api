package models

//model film
//utilisation de gorm pour les drivers postgres 
//ajout du tag cle primaire sur l'ID
type Film struct {
    Id int `json:"id" gorm:primaryKey`
    Title  string `json:"title"`
    Director string `json:"director"`
    Synopsis  string `json:"synopsis"`
    Score int `json:"score"`
    UserId int `json:"user"`
}