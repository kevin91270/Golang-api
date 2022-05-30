package models

//model user
//utilisation de gorm pour les drivers postgres 
//ajout du tag cle primaire sur l'ID
type User struct {
    Id int `json:"id" gorm:"primaryKey"`
    Pseudo  string `json:"pseudo"`
	Films []Film `gorm:"ForeignKey:user_id"`
}