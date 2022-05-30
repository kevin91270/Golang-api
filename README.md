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
│    │   ├── Add.go
│    │   ├── Delete.go
│    │   ├── GetAll.go
│    │   ├── Get.go
│    │   ├── handlers.go
│    │   └── Update.go
│    └── models
│        ├── user.go
│        └── film.go
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

### Routes
Film :
- `GET /films`
- `GET /films/:id`
- `POST /films {title, author, desc}`
- `PUT /films/:id {title, author, desc}`
- `DELETE /films/:id`

User :
- `GET /users`
- `GET /users/:id`
- `POST /users {title, author, desc}`
- `PUT /users/:id {title, author, desc}`
- `DELETE /users/:id`

## Modele

### Film :
- Id
- Title
- Director
- Synopsis
- Score
- UserId

### User :
- Id
- Pseudo
- Films