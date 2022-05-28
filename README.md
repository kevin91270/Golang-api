# Golang-api

API CRUD en GO pour le cours : API - Go language & concurrent computing
bibliotheque de film

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