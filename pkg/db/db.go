package db

import (
    "log"
    "github.com/kevin91270/Golang-api/pkg/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init() *gorm.DB {
	//ouverture de la connexion à la db
    dbURL := "postgres://pg:pass@localhost:5432/crud"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

	//migrer notre models
    db.AutoMigrate(&models.Film{})
	db.AutoMigrate(&models.User{})

    return db
}