package server

import (
	"fmt"
	"graficaconsola"
	"net/http"
	"os"
	"router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//StartServer inicia el servidor
func StartServer() {
	r := mux.NewRouter()

	// Rutas
	router.Routers(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	// Coors
	handler := cors.Default().Handler(r)

	// Grafica de consola
	graficaconsola.GraficaConsola()

	// Servidor encendido
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Println("Error", err)
	}
}
