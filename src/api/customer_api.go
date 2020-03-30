package api

import (
	"config"
	"models"
	"net/http"
	"tools"
)

//FindAllCustomer Muestra todos los clientes
func FindAllCustomer(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDatabase()
	if err != nil {
		tools.RespondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		customersModel := models.CustomersModel{
			Db: db,
		}
		customers, err2 := customersModel.FindAllCustomer()
		if err2 != nil {
			tools.RespondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			tools.RespondWithjson(response, http.StatusOK, customers)
		}
	}
}
