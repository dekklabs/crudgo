package router

import (
	"api"

	"github.com/gorilla/mux"
)

//Routers Configuración de todos las rutas
func Routers(router *mux.Router) {
	// Rutas Productos
	router.HandleFunc("/productos/findall", api.FindAllProduct).Methods("GET")

	// Rutas Proveedores
	router.HandleFunc("/suppliers/findall", api.FindAllSuppliers).Methods("GET")

	// Rutas Customers
	router.HandleFunc("/customers/findall", api.FindAllCustomer).Methods("GET")
}
