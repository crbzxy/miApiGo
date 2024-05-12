package api

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/api/users", UsersHandler).Methods("GET", "POST")
    router.HandleFunc("/api/users/{id}", UserHandler).Methods("GET", "PUT", "DELETE")
    return router
}
