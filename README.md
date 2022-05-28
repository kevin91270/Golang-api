# Golang-api

API CRUD en GO pour le cours : API - Go language & concurrent computing.  
Utilisation de Gorilla/Mux pour le routage plutôt que gin afin d'essayer d'autre methode
que celle vu en cours.  
Bibliotheque de films style PLEX, NETFLIX etc..

## STARTUP
Lancer la commande pour docker  
`docker compose up -d`

Lancer le projet  
`go run main.go`

## architecture
```
├── pkg
│    ├── db
│    │   ├── db.go
│    ├── handlers
│    │   ├── AddFilm.go
│    │   ├── DeleteFilm.go
│    │   ├── GetAllFilms.go
│    │   ├── GetFilm.go
│    │   ├── handlers.go
│    │   └── UpdateFilm.go
│    └── models
│        └── film.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

### Routes
Film :
- `GET /films`
- `GET /film/:id`
- `POST /film {title, author, desc}`
- `PUT /film/:id {title, author, desc}`
- `DELETE /users/:id`

## Modele

### Film :
- Id
- Title
- Author
- Desc