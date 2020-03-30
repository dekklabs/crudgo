package api

import (
	"config"
	"models"
	"net/http"
	"tools"
)

//FindAllSuppliers trae todos los proveedores
func FindAllSuppliers(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDatabase()
	if err != nil {
		tools.RespondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		suppliersConexion := models.SuppliersConexion{
			Db: db,
		}
		suppliers, err2 := suppliersConexion.FindAllModelSupplier()
		if err2 != nil {
			tools.RespondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			tools.RespondWithjson(response, http.StatusOK, suppliers)
		}
	}
}
