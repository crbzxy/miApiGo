package main

import (
	"log"
	"miapigo/internal/api"
	"net/http"
)

func main() {
	router := api.NewRouter() // Configuración de rutas
	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
